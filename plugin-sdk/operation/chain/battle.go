package chain

import (
	"context"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/callback"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/battle"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/chain/spi"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/chain/spi/config"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/pet"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/runner"
	"strings"
	"time"
)

var BattleStateMachineInstance = &BattleStateMachine{}
var battleCtx context.Context
var battleCancel context.CancelFunc

type BattleStateMachine struct{}

func (BattleStateMachine) HandleFinishEvent() {
	battleCancel()
}

func (BattleStateMachine) HandleStartEvent() {
	battleCtx, battleCancel = context.WithCancel(context.Background())
	go startBattleTask(battleCtx, battleCancel)
	plugin_log.Info("开启挂机 ")
}

func startBattleTask(ctx context.Context, cancel context.CancelFunc) {
	reporter := biz_callback.NewDataReporter()
	callbackInterface := &spi.BattleReportCallback{}
	reporter.Start(callbackInterface)

	currentUser := status.GetLoginUser()
	battleConfig := config.GetBattleConfig(currentUser.LoginName)

	pet.PetServiceInstance.SaveUnBattlePet()
	// 初始化PetId
	if battleConfig.PetId == "" {
		// 获取身上的宠物列表
		list := pet.PetServiceInstance.GetAllPets()
		if battleConfig.PetName == "" {
			battlePet := battle.GetBattlePet(list)
			battleConfig.PetId = battlePet.Id
		}
		if battleConfig.PetId == "" {
			for _, battlePet := range list {
				if strings.Contains(battlePet.Name, battleConfig.PetName) {
					battleConfig.PetId = battlePet.Id
					break
				}
			}
		}
	}

	// 初始化技能ID
	if battleConfig.SkillId == "" {
		SkillList, _ := pet.PetServiceInstance.GetPetSkillList(battleConfig.PetId)
		for _, Skill := range SkillList {
			if strings.Contains(Skill.SkillName, battleConfig.SkillName) {
				battleConfig.SkillId = Skill.SkillId
				break
			}
		}
		if battleConfig.SkillId == "" {
			battleConfig.SkillId = "1"
		}
	}
	pet.PetServiceInstance.SetBattlePet(battleConfig.PetId)
	battleConfig2 := model.BattleConfig{
		PetId:              battleConfig.PetId,
		SkillId:            battleConfig.SkillId,
		MapId:              battleConfig.MapId,
		Difficulty:         battleConfig.Difficulty,
		SkipMonsters:       battleConfig.SkipMonsters,
		CatchPets:          battleConfig.CatchPets,
		RunWhenCatchFailed: battleConfig.RunWhenCatchFailed,
		RunWhenNotCatch:    battleConfig.RunWhenNotCatch,
		Balls:              battleConfig.Balls,
		Rubbish:            battleConfig.Rubbish,
		CatchHpThreshold:   battleConfig.CatchHpThreshold,
		SaveAfterCatch:     battleConfig.SaveAfterCatch,
	}

	for {
		select {
		case <-ctx.Done():
			plugin_log.Info("停止挂机任务")
			runner.StopTask()
			return
		default:
			result := battle.BattleServiceImplInstance.FightOneTime(battleConfig2)
			reporter.SendData(result)
			time.Sleep(1 * time.Second)
		}
	}
}
