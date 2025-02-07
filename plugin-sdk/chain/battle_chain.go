package chains

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain/config"
	"strings"
)

var BattleReportCallbackInstance = &BattleReportCallback{}

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
			status.SetBattleStatus(status.Waiting2Stop)
			return
		}
	}
	if config.FinishConfig.CatchPet > 0 {
		if config.FinishConfig.CatchPet <= CatchCount {
			plugin_log.Info("捕捉数量达到预设值 %d, 退出战斗", config.FinishConfig.CatchPet)
			status.SetBattleStatus(status.Waiting2Stop)
			return
		}
	}

}

func (*BattleReportCallback) StopCallback() {
	BattleCount = 0
	CatchCount = 0
}

func BattleChain() {
	currentUser := status.GetLoginUser()
	config := config.GetBattleConfig(currentUser.LoginName)
	plugin_sdk.PetServiceInstance.SaveUnBattlePet()
	// 初始化PetId
	if config.PetId == "" {
		// 获取身上的宠物列表
		list := plugin_sdk.PetServiceInstance.GetAllPets()
		if config.PetName == "" {
			battlePet := getBattlePet(list)
			config.PetId = battlePet.Id
		}
		if config.PetId == "" {
			for _, pet := range list {
				if strings.Contains(pet.Name, config.PetName) {
					config.PetId = pet.Id
					break
				}
			}
		}
	}

	// 初始化技能ID
	if config.SkillId == "" {
		SkillList, _ := plugin_sdk.PetServiceInstance.GetPetSkillList(config.PetId)
		for _, Skill := range SkillList {
			if strings.Contains(Skill.SkillName, config.SkillName) {
				config.SkillId = Skill.SkillId
				break
			}
		}
		if config.SkillId == "" {
			config.SkillId = "1"
		}
	}
	plugin_sdk.PetServiceInstance.SetBattlePet(config.PetId)
	battleConfig := model.BattleConfig{
		PetId:              config.PetId,
		SkillId:            config.SkillId,
		MapId:              config.MapId,
		Difficulty:         config.Difficulty,
		SkipMonsters:       config.SkipMonsters,
		CatchPets:          config.CatchPets,
		RunWhenCatchFailed: config.RunWhenCatchFailed,
		RunWhenNotCatch:    config.RunWhenNotCatch,
		Balls:              config.Balls,
		Rubbish:            config.Rubbish,
		CatchHpThreshold:   config.CatchHpThreshold,
		SaveAfterCatch:     config.SaveAfterCatch,
	}
	plugin_sdk.BattleServiceImplInstance.FightByConfig(battleConfig, BattleReportCallbackInstance)
}

func getBattlePet(pets []*model.Pet) *model.Pet {
	for _, pet := range pets {
		if pet.IsBattle {
			return pet
		}
	}
	return nil
}
