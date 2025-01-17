package util

import (
	"bytes"
	"encoding/json"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
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
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
