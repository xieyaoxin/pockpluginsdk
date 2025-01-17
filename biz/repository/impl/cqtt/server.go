package cqtt

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"plugin-sdk/biz/log"
	"plugin-sdk/biz/model"
	"plugin-sdk/biz/status"
	util "plugin-sdk/biz/utils"
	"strings"
)

//const endpoint = "http://124.223.98.176:1000"

const endpoint = "https://101357.xyz:83"

func InitParam() map[string]string {
	return make(map[string]string)
}

func callServerFormInterface(interfaceName string, params map[string]string) string {
	if status.IsBattleParsing() {
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
	req.Header.Add("Cookie", "PHPSESSID="+status.GetLoginToken())
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
	if status.IsBattleParsing() {
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
	req.Header.Add("Cookie", "PHPSESSID="+status.GetLoginToken())
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
		Login(*status.GetLoginUser(), 0)
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
	log.Info("start login ")
	retryTimes = retryTimes + 1

	//signature := callLogin()
	//if signature == "" {
	//
	//	return Login(user, retryTimes)
	//}
	params := make(map[string]string)
	params["username"] = user.LoginName
	//params["mac"] = strconv.FormatInt((time.Now().UnixNano() / 1e6), 10)
	//params["sign"] = signature
	params["password"] = user.Password
	dealPc := callServerFormInterface("/passport/dealPc.php", params)
	if dealPc == "" {
		return Login(user, retryTimes)
	}
	loginResult := CallServerGetInterface("/login/login.php", util.InitParam())
	if strings.Contains(loginResult, "/passport/login.php") {
		// todo 错误原因解析
		log.Info(loginResult)
		return ""
	}
	CallServerGetInterface("index.php", InitParam())
	return user.TempToken
}

// 访问登录页面 获取splite
func callLogin() string {
	responseHtml := CallServerGetInterface("/passport/login.php", InitParam())
	resultLines := strings.Split(responseHtml, "\n")
	for lineNumber := range resultLines {
		line := resultLines[lineNumber]
		if strings.Contains(line, "var s=") {
			return strings.Replace(strings.Split(line, "=")[1], ";", "", 1)
		}
	}
	return ""
}

func GetSessionId(userName string) string {
	return MD5(userName)
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func generateRandomSeed() string {
	length := 10
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
