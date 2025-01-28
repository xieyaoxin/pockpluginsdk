package plugin_sdk

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"

	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	status2 "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	biz_callback "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/callback"
	"strings"
	"time"
)

// 天梯
var TtServiceImplInstance = &ttServiceImpl{}
var instance = repository.GetTtRepository()
var battleRepositoryInstance = repository.GetBattleRepository()

type ttServiceImpl struct {
}

// todo 通天/副本/挂机 互斥
func (*ttServiceImpl) StartTtLoop(config *model.TtConfig, callbackInterface *biz_callback.TtReportCallbackInterface) bool {

	// 后续加锁
	if status2.GetConflictTask() {
		return false
	}
	status2.SetBattleStatus(status2.Running)
	reporter := biz_callback.NewDataReporter()
	reporter.Start(callbackInterface)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				status2.SetTtBattleStatus(status2.NotReady)
				time.Sleep(time.Second)
				reporter.Stop(callbackInterface)
			}
		}()
		err := PetServiceInstance.SaveUnBattlePet()
		if err != nil {
			return
		}
		for {
			err = fight4TtOnce(config)
			if err != nil {
				plugin_log.Error(err.Error())
				if err.Error() == "战斗失败" && config.LoopTt {
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
func fight4TtOnce(config *model.TtConfig) error {
	CurrentLevel := instance.EnterTt()
	plugin_log.Info("当前层数： %s", CurrentLevel)
	// todo 判断当前层
	// 判断是否需要花费水晶
	checkResult := instance.ShouldPaySj("")
	// 需要花费水晶开启天梯
	if checkResult != "b" {
		// 花费水晶打开天梯
		checkResult = instance.ShouldPaySj("do")
		for checkResult == "c" {
			err := ArticleServiceInstance.UseSjk()
			if err != nil {
				return err
			}
			// 再次尝试
			checkResult = instance.ShouldPaySj("do")
		}
	}
	monster, err := enterTTMap(config)
	if err != nil {
		return err
	}
	battleConfig := model.BattleConfig{
		PetId:   config.PetId,
		SkillId: config.SkillId,
	}
	result := fight(battleConfig, monster)
	if !result {
		return errors.New("战斗失败")
	}
	return nil
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

		r1 := instance.Pay30SJ(config.PetId)
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
