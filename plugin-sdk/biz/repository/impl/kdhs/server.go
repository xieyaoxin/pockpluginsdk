package KDHS

import (
	"encoding/json"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	status2 "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//const endpoint = "http://124.223.98.176:1000"

const endpoint = "http://43.248.129.148:567"

func InitParam() map[string]string {
	return make(map[string]string)
}

func callServerFormInterface(interfaceName string, params map[string]string) string {
	if status2.IsBattleParsing() {
		panic("正在停止战斗任务")
	}
	payload := url.Values{}
	for k, v := range params {
		payload.Set(k, v)
	}

	req, err := http.NewRequest(http.MethodPost,
		endpoint+interfaceName,
		strings.NewReader(payload.Encode()))
	if err != nil {
		return ""
	}

	req.Header.Add("Content-Type",
		"application/x-www-form-urlencoded; param=value")
	req.Header.Add("Accept", "text/javascript, text/html, application/xml, text/xml, */*")
	req.Header.Add("X-Prototype-Version", "1.5.0")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Cookie", "PHPSESSID="+status2.GetLoginToken())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// retry
		return callServerFormInterface(interfaceName, params)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	utf8, _ := util.GbkToUtf8(data)
	return string(utf8)
}

// 调用口袋接口,
func CallServerGetInterface(interfaceName string, param map[string]string) string {
	if status2.IsBattleParsing() {
		panic("正在停止战斗任务")
	}
	// 设置统一延时 300ms
	client := &http.Client{}
	if strings.HasSuffix(endpoint, "/") && strings.HasPrefix(interfaceName, "/") {
		interfaceName = interfaceName[1:]
	}
	if !strings.HasSuffix(endpoint, "/") && !strings.HasPrefix(interfaceName, "/") {
		interfaceName = "/" + interfaceName
	}
	apiURL := endpoint + interfaceName
	req, err := http.NewRequest("GET", apiURL, nil)
	//添加查询参数
	query := req.URL.Query()
	for k, v := range param {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()
	//log.Info(req.URL.String())
	req.Header.Add("Accept", "text/javascript, text/html, application/xml, text/xml, */*")
	req.Header.Add("X-Prototype-Version", "1.5.0")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Cookie", "PHPSESSID="+status2.GetLoginToken())
	if err != nil {
		log.Info("构造请求失败, err:%v\n\n", err)
		return CallServerGetInterface(interfaceName, param)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Info("调用接口错误,错误原因:%v\n\n", err)
		return CallServerGetInterface(interfaceName, param)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("get resp failed,err:%v\n\n", err)
		return CallServerGetInterface(interfaceName, param)
	}
	result := string(b)
	if resp.StatusCode != 200 {
		log.Info("调用接口异常,重试  接口返回状态码 %d 返回内容 ： %s", resp.StatusCode, string(b))
	}

	if strings.Contains(result, "登录") && interfaceName != "/passport/login.php" {
		log.Info("重新登录： uri: %s param: %s, 原因 %s", interfaceName, util.MapToJsonString(param), result)
		Login(*status2.GetLoginUser(), 0)
		return CallServerGetInterface(interfaceName, param)
	}

	if strings.EqualFold(result, "数据错误1！") {
		log.Info("数据异常. 重新调用方法 uri: %s param: %s ", interfaceName, util.MapToJsonString(param))
		return CallServerGetInterface(interfaceName, param)
	}
	utf8, _ := util.GbkToUtf8(b)
	return string(utf8)
}

// 调用登录接口  返回session
func Login(user model.User, retryTimes int) string {
	if retryTimes > 5 {
		log.Error("登录重试次数大于5,停止登录")
		return ""
	}
	log.Info("开始登陆 ")
	retryTimes = retryTimes + 1
	signature := callLogin()
	//signature := callLogin()
	//if signature == "" {
	//
	//	return Login(user, retryTimes)
	//}
	params := make(map[string]string)
	params["username"] = user.LoginName
	//params["mac"] = strconv.FormatInt((time.Now().UnixNano() / 1e6), 10)
	params["sign"] = signature
	params["password"] = user.Password

	loginResult := callServerFormInterface("/passport/login_ac.php", params)
	var loginResponse map[string]interface{}
	json.Unmarshal([]byte(loginResult), &loginResponse)
	if loginResponse["code"] == nil || int(loginResponse["code"].(float64)) != 0 {
		return Login(user, retryTimes)
	}
	log.Info("登陆成功 ")
	return user.TempToken
}

// 访问登录页面 获取splite
func callLogin() string {
	responseHtml := CallServerGetInterface("/passport/login.php", InitParam())
	resultLines := strings.Split(responseHtml, "\n")
	for lineNumber := range resultLines {
		line := resultLines[lineNumber]
		if strings.Contains(line, "<input type=\"hidden\" id=\"sign\" ") {
			return strings.Split(strings.Replace(strings.Split(strings.Trim(line, " "), " ")[3], "\"", " ", -1), " ")[1]
		}
	}
	return ""
}

func GetSessionId(userName string) string {
	return util.MD5(userName)
}
