package config

import (
	"encoding/json"
	"fmt"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain/chain_model"
	"io"
	"os"
)

func GetBattleConfig(AccountName string) *chain_model.BattleChainConfig {
	configFile := fmt.Sprintf("config/%s/挂机配置.json", AccountName)
	configString := ReadFromFile(configFile)
	var config = &chain_model.BattleChainConfig{}
	err := json.Unmarshal(configString, config)
	if err != nil {
		panic("获取战斗配置失败，账号: " + AccountName)
	}
	return config
}

func GetFusionConfig(AccountName string) *chain_model.FusionChainConfig {
	configFile := fmt.Sprintf("config/%s/合神配置.json", AccountName)
	configString := ReadFromFile(configFile)
	var config = &chain_model.FusionChainConfig{}
	err := json.Unmarshal(configString, config)
	if err != nil {
		panic("获取合神配置失败，账号: " + AccountName)
	}
	return config
}

func GetChainConfig(AccountName string) []string {
	configFile := fmt.Sprintf("config/%s/全局配置.txt", AccountName)
	configString := ReadFromFile(configFile)
	var config = []string{}
	err := json.Unmarshal(configString, &config)
	if err != nil {
		panic("获取全局配置失败，账号: " + AccountName)
	}
	return config
}

func ReadFromFile(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		// 打开文件失败
		return []byte("{}")
	}
	var data []byte
	buf := make([]byte, 1024)
	for {
		// 将文件中读取的byte存储到buf中
		n, err1 := f.Read(buf)
		if err1 != nil && err1 != io.EOF {
			plugin_log.Fatal(err1.Error())
		}
		if n == 0 {
			break
		}
		// 将读取到的结果追加到data切片中
		data = append(data, buf[:n]...)
	}
	return data
}

func GetNirvanaConfig(AccountName string) *chain_model.NirvanaChainConfig {
	configFile := fmt.Sprintf("config/%s/涅槃配置.json", AccountName)
	configString := ReadFromFile(configFile)
	var config = &chain_model.NirvanaChainConfig{}
	err := json.Unmarshal(configString, config)
	if err != nil {
		panic("获取涅槃配置失败，账号: " + AccountName)
	}
	return config
}
