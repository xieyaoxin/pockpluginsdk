package service

import (
	"encoding/json"
	"plugin-sdk/biz/log"
	"plugin-sdk/biz/model"
	"plugin-sdk/biz/service/impl"
	"plugin-sdk/biz/status"
	"time"
)

func (app *App) StartBattle(battleConfig string) string {
	BattleConfig := &model.BattleConfig{}
	err := json.Unmarshal([]byte(battleConfig), BattleConfig)
	if err != nil {
		log.Info("开始战斗的配置为： %s", battleConfig)
		return buildSuccessResponse("解析json失败")
	}
	result := impl.BattleServiceImpl.Fight(*BattleConfig)
	// todo init config
	if result {
		return buildSuccessResponse("开启战斗成功")
	} else {
		return buildFailedResponse("存在进行中的任务,无法启动")
	}
}

func (app *App) GetBattleStatus() string {
	return buildSuccessResponse(status.FightStatus.BattleStatus)
}

func (app *App) StopBattle() string {
	status.SetBattleStatus(status.Waiting2Stop)
	for {
		time.Sleep(1 * time.Second)
		BattleStatus := status.IsBattleRunning()
		if !BattleStatus {
			break
		}
	}
	return buildSuccessResponse(status.FightStatus.BattleStatus == status.NotReady)
}
