package battle

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_sdk_const"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	biz_callback "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/callback"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/article"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/pet"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/runner"
	"strings"
	"time"
)

var TtServiceImplInstance = &ttServiceImpl{}
var battleRepository = repository.GetBattleRepository()
var BattleServiceImplInstance = &battleService{}

type ttServiceImpl struct {
}

func (inst *ttServiceImpl) Start() {
	err := runner.StartTask(runner.BATTLE)
	if err != nil {
		return
	}
}

type battleService struct {
}

func (inst *battleService) FightByConfig(BattleConfig model.BattleConfig, callbackInterface biz_callback.BattleReportCallbackInterface) bool {

	status.SetTaskType(status.BATTLE)
	reporter := biz_callback.NewDataReporter()
	reporter.Start(callbackInterface)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				status.SetTaskType(status.NONE)
				time.Sleep(time.Second)
				reporter.Stop(callbackInterface)
			}
		}()
		err := pet.PetServiceInstance.SaveUnBattlePet()
		if err != nil {
			return
		}
		for {
			result := inst.FightOneTime(BattleConfig)
			reporter.SendData(result)
			time.Sleep(1 * time.Second)
		}
	}()
	return true
}

// FightOneTime 普通地图 根据配置捕捉或击杀 - 完成一次进入地图的战斗
// 进入地图失败时 返回false
func (inst *battleService) FightOneTime(BattleConfig model.BattleConfig) string {

	monster, err := battleRepository.SelectAndEnterMap(BattleConfig.MapId, BattleConfig.PetId, BattleConfig.Difficulty)
	if err != nil {
		plugin_log.Error("进入地图失败")
		return "进入地图异常"
	}
	for {
		// 00: 不在捕捉范围内  01: 捕捉失败 11: 捕捉成功; 10: 战斗失败 / 战斗成功
		result := catchPet(BattleConfig, monster)
		switch result {
		case "11":
			return "捕捉成功"
		case "10":
			return "战斗结束"
		case "00":
			if BattleConfig.RunWhenNotCatch {
				plugin_log.Info("当前怪物不在捕捉列表中,跳过")
			} else {
				inst.Fight(BattleConfig, monster)
			}
			return "不在捕捉范围内"
		case "01":
			if BattleConfig.RunWhenCatchFailed {
				plugin_log.Info("捕捉失败,跳过本次战斗")
				return "捕捉失败"
			} else {
				inst.Fight(BattleConfig, monster)
				return "战斗结束"
			}
		}

	}
}

func (inst *battleService) Fight(BattleConfig model.BattleConfig, monster *model.Monster) bool {
	for {
		result := battleRepository.FightOnce(BattleConfig.SkillId, monster)
		if result == "10" {
			time.Sleep(time.Duration(2000) * time.Millisecond)
		} else {
			return result == "11"
		}
	}
}

// catchPet: 00: 不在捕捉范围内  01: 捕捉失败 11: 捕捉成功; 10: 战斗失败 / 战斗成功
func catchPet(BattleConfig model.BattleConfig, monster *model.Monster) string {
	NeedCatch := false
	for _, CatchMonsterName := range BattleConfig.CatchPets {
		if strings.Contains(monster.Name, CatchMonsterName) {
			NeedCatch = true
			break
		}
	}
	if !NeedCatch {
		return "00"
	}
	for int(monster.CurrentHpRate) > BattleConfig.CatchHpThreshold {
		result := battleRepository.FightOnce(BattleConfig.SkillId, monster)
		// 战斗成功 / 战斗失败 -> 返回捕捉失败
		if result == "00" || result == "11" {
			return "10"
		}
		if int(monster.CurrentHpRate) >= BattleConfig.CatchHpThreshold {
			break
		}
		time.Sleep(time.Duration(2000) * time.Millisecond)
	}
	BallList := getBallNameListByMonsterName(monster.Name, BattleConfig.Balls)
	if len(BallList) > 0 {
		BallId := BallList[0].ID
		plugin_log.Info("开始捕捉 %s, 使用 %s 球", monster.Name, BallList[0].Name)
		result := battleRepository.CatchPet(monster, BallId)
		if result {
			if BattleConfig.SaveAfterCatch {
				err := pet.PetServiceInstance.SaveUnBattlePet()
				if err != nil {
					return ""
				}
			}

			return "11"
		} else {
			return "01"
		}
	} else {
		plugin_log.Error("找不到对应精灵球")
		return "01"
	}
}

func getBallNameListByMonsterName(monsterName string, balls []string) []*model.Article {
	// 部分
	BallNameList := plugin_sdk_const.GetBallByMonster(monsterName, balls)
	ballList, _ := article.ArticleServiceInstance.QueryArticleListByNameLists(BallNameList)

	return ballList
}

// InitCatchBmConfig 提供一个抓BM的配置
func InitCatchBmConfig() *model.BattleConfig {
	Pet := pet.PetServiceInstance.GetBattlePet()
	return &model.BattleConfig{
		PetId:              Pet.Id,
		SkillId:            "1",
		MapId:              "1",
		Difficulty:         "1",
		SkipMonsters:       []string{},
		CatchPets:          []string{"波姆"},
		RunWhenCatchFailed: true,
		RunWhenNotCatch:    true,
		Balls:              []string{},
		Rubbish:            []string{},
		CatchHpThreshold:   100,
	}
}

func GetBattlePet(pets []*model.Pet) *model.Pet {
	for _, pet := range pets {
		if pet.IsBattle {
			return pet
		}
	}
	return nil
}

func InitBattlePet(PetId, PetName, SkillId, SkillName string) (string, string) {
	pet.PetServiceInstance.SaveUnBattlePet()
	// 初始化PetId
	if PetId == "" {
		// 获取身上的宠物列表
		list := pet.PetServiceInstance.GetAllPets()
		if PetName == "" {
			battlePet := GetBattlePet(list)
			PetId = battlePet.Id
		}
		if PetId == "" {
			for _, pet := range list {
				if strings.Contains(pet.Name, PetName) {
					PetId = pet.Id
					break
				}
			}
		}
	}

	// 初始化技能ID
	if SkillId == "" {
		SkillList, _ := pet.PetServiceInstance.GetPetSkillList(PetId)
		for _, Skill := range SkillList {
			if strings.Contains(Skill.SkillName, SkillName) {
				SkillId = Skill.SkillId
				break
			}
		}
		if SkillId == "" {
			SkillId = "1"
		}
	}
	pet.PetServiceInstance.SetBattlePet(PetId)
	return PetId, SkillId
}
