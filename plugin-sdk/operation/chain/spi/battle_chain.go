package spi

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/chain/spi/config"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/runner"
)

var BattleCount = 0
var CatchCount = 0

type BattleReportCallback struct {
}

func (*BattleReportCallback) Callback(input interface{}) {
	BattleCount++
	if input == "捕捉成功" {
		CatchCount++
	}
	plugin_log.Info("当前为第 %d 次战斗", BattleCount)
	currentUser := status.GetLoginUser()
	// 获取当前用户操作配置
	config := config.GetBattleConfig(currentUser.LoginName)
	if config.FinishConfig == nil {
		return
	}
	if config.FinishConfig.BattleCount > 0 {
		if config.FinishConfig.BattleCount <= BattleCount {
			plugin_log.Info("战斗次数达到预设值 %d, 退出战斗", config.FinishConfig.BattleCount)
			runner.StopTask()
			return
		}
	}
	if config.FinishConfig.CatchPet > 0 {
		if config.FinishConfig.CatchPet <= CatchCount {
			plugin_log.Info("捕捉数量达到预设值 %d, 退出战斗", config.FinishConfig.CatchPet)
			runner.StopTask()
			return
		}
	}

}

func (*BattleReportCallback) StopCallback() {
	BattleCount = 0
	CatchCount = 0
}
