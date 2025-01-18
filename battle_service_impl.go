package plugin_sdk

import (
	"github.com/xieyaoxin/plugin-sdk/biz/log"
	"github.com/xieyaoxin/plugin-sdk/biz/model"
	"github.com/xieyaoxin/plugin-sdk/biz/repository"
	"github.com/xieyaoxin/plugin-sdk/biz/status"
	"strings"
	"time"
)

var impl = repository.GetBattleRepository()
var BattleServiceImpl = &battleService{}

type battleService struct {
}

func (instance *battleService) Fight(BattleConfig model.BattleConfig) bool {
	// 后续加锁
	if status.GetConflictTask() {
		return false
	}
	status.SetBattleStatus(status.Running)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				status.SetBattleStatus(status.NotReady)
			}
		}()
		err := PetServiceInstance.SaveUnBattlePet()
		if err != nil {
			return
		}
		for {
			FightOneTime(BattleConfig)
		}
	}()
	return true
}

// FightOneTime 普通地图 根据配置捕捉或击杀 - 完成一次进入地图的战斗
// 进入地图失败时 返回false
func FightOneTime(BattleConfig model.BattleConfig) bool {

	monster, err := impl.SelectAndEnterMap(BattleConfig.MapId, BattleConfig.PetId)
	if err != nil {
		log.Error("进入地图失败")
		return false
	}
	for {
		result := catchPet(BattleConfig, monster)
		switch result {
		case "11":
			return true
		case "10":
			return true
		case "00":
			if BattleConfig.RunWhenNotCatch {
				log.Info("当前怪物不在捕捉列表中,跳过")
			} else {
				fight(BattleConfig, monster)
			}
			return true
		case "01":
			if BattleConfig.RunWhenCatchFailed {
				log.Info("捕捉失败,跳过本次战斗")
			} else {
				fight(BattleConfig, monster)
			}
			return true
		}

	}
}

func fight(BattleConfig model.BattleConfig, monster *model.Monster) {
	for {
		result := impl.FightOnce(BattleConfig.SkillId, monster)
		if result != "10" {
			time.Sleep(time.Duration(2000) * time.Millisecond)
			return
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
	for monster.CurrentHpRate > BattleConfig.CatchHpThreshold {
		result := impl.FightOnce(BattleConfig.SkillId, monster)
		// 战斗成功 / 战斗失败 -> 返回捕捉失败
		if result == "00" || result == "11" {
			return "10"
		}
		if monster.CurrentHpRate >= BattleConfig.CatchHpThreshold {
			break
		}
		time.Sleep(time.Duration(2000) * time.Millisecond)
	}
	BallList := getBallNameListByMonsterName(monster.Name)
	if len(BallList) > 0 {
		BallId := BallList[0].ID
		result := impl.CatchPet(monster, BallId)
		if result {
			err := PetServiceInstance.SaveUnBattlePet()
			if err != nil {
				return ""
			}
			return "11"
		} else {
			return "01"
		}
	} else {
		log.Error("找不到对应精灵球")
		return "01"
	}
}

func getBallNameListByMonsterName(monsterName string) []*model.Article {
	ballName := monsterName + "·精灵球"
	ballList, _ := ArticleServiceInstance.QueryArticleList(ballName)
	return ballList
}
