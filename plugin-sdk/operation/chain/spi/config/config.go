package config

import (
	"encoding/json"
	"fmt"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	chain_model2 "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/chain/spi/chain_model"
	"io"
	"os"
)

func GetBattleConfig(AccountName string) *chain_model2.BattleChainConfig {
	configFile := fmt.Sprintf("config/%s/挂机配置.json", AccountName)
	configString := ReadFromFile(configFile)
	var config = &chain_model2.BattleChainConfig{}
	err := json.Unmarshal(configString, config)
	if err != nil {
		panic("获取战斗配置失败，账号: " + AccountName)
	}
	return config
}

func GetFusionConfig(AccountName string) *chain_model2.FusionChainConfig {
	configFile := fmt.Sprintf("config/%s/合神配置.json", AccountName)
	configString := ReadFromFile(configFile)
	var config = &chain_model2.FusionChainConfig{}
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

func GetTtConfig(AccountName string) *chain_model2.TtChainConfig {
	configFile := fmt.Sprintf("config/%s/通天配置.json", AccountName)
	configString := ReadFromFile(configFile)
	var config = &chain_model2.TtChainConfig{}
	err := json.Unmarshal(configString, config)
	if err != nil {
		panic("获取通天配置失败，账号: " + AccountName)
	}
	return config
}

func GetDungeonInstanceConfig(AccountName string) *chain_model2.DungeonChainConfig {
	configFile := fmt.Sprintf("config/%s/副本配置.json", AccountName)
	configString := ReadFromFile(configFile)
	var config = &chain_model2.DungeonChainConfig{}
	err := json.Unmarshal(configString, config)
	if err != nil {
		panic("获取副本配置失败，账号: " + AccountName)
	}
	return config
}

func GetDefaultTimerConfig(AccountName string) []*model.TimerConfig {
	configFile := fmt.Sprintf("config/%s/定时配置.json", AccountName)
	configString := ReadFromFile(configFile)
	var config = []*model.TimerConfig{}
	err := json.Unmarshal(configString, &config)
	if err != nil {
		panic("获取定时任务配置失败，账号: " + AccountName)
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

func GetNirvanaConfig(AccountName string) *chain_model2.NirvanaChainConfig {
	configFile := fmt.Sprintf("config/%s/涅槃配置.json", AccountName)
	configString := ReadFromFile(configFile)
	var config = &chain_model2.NirvanaChainConfig{}
	err := json.Unmarshal(configString, config)
	if err != nil {
		panic("获取涅槃配置失败，账号: " + AccountName)
	}
	return config
}
