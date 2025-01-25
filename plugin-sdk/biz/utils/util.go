package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	status2 "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"log"
	"time"
)

func InSlice(slice []*interface{}, target interface{}) bool {
	for _, item := range slice {
		if target == item {
			return true
		}
	}
	return false
}

func InStringSlice(slice []string, target string) bool {
	for _, item := range slice {
		if target == item {
			return true
		}
	}
	return false
}

func Json2Obj(jsonString string, obj interface{}) interface{} {
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		log.Printf("解析失败： %s", jsonString)
	}
	return obj
}

func MapToJsonString(param any) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func ListToJson(list any) string {
	result, _ := json.Marshal(list)
	return string(result)
}

func String2JsonArray(jsonArrayString string) []any {
	var result []any
	err := json.Unmarshal([]byte(jsonArrayString), &result)
	if err != nil {
		log.Printf("解析失败： %s", jsonArrayString)
	}
	return result
}

func Json2Map(jsonArrayString string) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonArrayString), &result)
	if err != nil {
		log.Printf("解析失败： %s", jsonArrayString)
	}
	return result
}

func InitParam() map[string]string {
	return make(map[string]string)
}

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	//reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	//d, e := ioutil.ReadAll(reader)
	//if e != nil {
	//	return nil, e
	//}
	return s, nil
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func GenerateRandomSeed() string {
	if status2.SESSION_ID_KEEP_TYPE {
		return ""
	}
	length := 10
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func CalculateTime(startTime time.Time) (time.Duration, time.Duration) {
	delta := time.Since(startTime)

	minute := delta / time.Minute
	sec := (delta - minute*time.Minute) / time.Second
	return minute, sec
}
