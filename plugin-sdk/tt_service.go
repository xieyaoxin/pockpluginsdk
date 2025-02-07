package plugin_sdk

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	biz_callback "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/callback"

	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	status2 "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"strings"
	"time"
)

// 天梯
var TtServiceImplInstance = &ttServiceImpl{}
var battleRepositoryInstance = repository.GetBattleRepository()

type ttServiceImpl struct {
}

// todo 通天/副本/挂机 互斥
func (*ttServiceImpl) StartTt(config *model.TtConfig, callbackInterface biz_callback.TtReportCallbackInterface) bool {

	// 后续加锁
	if status2.GetConflictTask() {
		plugin_log.Error("任务冲突,")
		return false
	}
	status2.SetTaskType(status2.TT)
	status2.SetBattleStatus(status2.Running)
	reporter := biz_callback.NewDataReporter()
	reporter.Start(callbackInterface)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				status2.SetBattleStatus(status2.NotReady)
				time.Sleep(time.Second)
				reporter.Stop(callbackInterface)
			}
		}()
		err := PetServiceInstance.SaveUnBattlePet()
		if err != nil {
			return
		}
		for {
			currentLevel, err1 := fight4TtOnce(config)
			reporter.SendData(currentLevel)
			if err1 != nil {
				plugin_log.Error(err1.Error())
				if err1.Error() == "战斗失败" && config.LoopTt {
					// 循环挂机
					continue
				}
				return
			}
		}
	}()
	return true
}

func (*ttServiceImpl) StopTtLoop() {

}
func fight4TtOnce(config *model.TtConfig) (string, error) {
	CurrentLevel := repository.GetTtRepository().EnterTt()
	plugin_log.Info("当前层数： %s", CurrentLevel)
	// todo 判断当前层
	// 判断是否需要花费水晶
	checkResult := repository.GetTtRepository().ShouldPaySj("")
	// 需要花费水晶开启天梯
	if checkResult != "b" {
		// 花费水晶打开天梯
		checkResult = repository.GetTtRepository().ShouldPaySj("do")
		for checkResult == "c" {
			err := ArticleServiceInstance.UseSjk()
			if err != nil {
				return CurrentLevel, err
			}
			// 再次尝试
			checkResult = repository.GetTtRepository().ShouldPaySj("do")
		}
	}
	monster, err := enterTTMap(config)
	if err != nil {
		return CurrentLevel, err
	}
	battleConfig := model.BattleConfig{
		PetId:   config.PetId,
		SkillId: config.SkillId,
	}
	result := BattleServiceImplInstance.Fight(battleConfig, monster)
	if !result {
		return CurrentLevel, errors.New("战斗失败")
	}
	return CurrentLevel, nil
}

func enterTTMap(config *model.TtConfig) (*model.Monster, error) {
	monster, err := battleRepositoryInstance.EnterMap(config.PetId)
	if err == nil {
		return monster, nil
	}
	result := err.Error()
	// 异常情况处理
	for strings.Contains(result, "继续31层，将收取200水晶，是否继续") {
		// todo 增加30层时候的判断

		r1 := repository.GetTtRepository().Pay30SJ(config.PetId)
		if !r1 {
			err = ArticleServiceInstance.UseSjk()
			if err != nil {
				return nil, err
			}
		}
		monster, err = battleRepositoryInstance.EnterMap(config.PetId)
		if err == nil {
			result = ""
		} else {
			plugin_log.Error("进入31层失败 失败原因 %s", err)
		}
	}
	return monster, err
}
