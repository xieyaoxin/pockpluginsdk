package plugin_sdk

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_sdk_const"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"time"
)

var DungeonInstanceServiceImplInstance = &dungeonInstanceServiceImpl{}
var repositoryInstance = repository.GetDungeonInstanceRepository()
var MaxBattleFailedTimes = 5

type dungeonInstanceServiceImpl struct {
}

func (*dungeonInstanceServiceImpl) FightDungeonOnce(Config *model.DungeonInstanceConfig) error {

	// 校验副本状态
	CurrentStage, CountDown := repositoryInstance.GetDungeonInstanceStatus(Config.MapId)
	plugin_log.Info("当前副本进度: %d, 当前副本倒计时 %d ", CurrentStage, CountDown)
	if CurrentStage == -1 || CountDown == -1 {
		return errors.New(plugin_sdk_const.GET_DUNGEON_INSTANCE_STATUS_FAILED)
	}
	if CountDown > 0 && CurrentStage == 1 {
		plugin_log.Info("副本未开启")
		// 倒计时大于0 当前关卡等于1 -> 副本已结束
		if Config.ForceFight == false {
			return errors.New(plugin_sdk_const.GET_DUNGEON_NOT_OPEN)
		} else {
			plugin_log.Info("消耗水晶开启副本")
			err := OpenDungeonInstance(Config)
			if err != nil {
				return err
			}
		}
	}

	for {

		// 进入地图
		monster, err := repositoryInstance.EnterMap(Config.PetId, Config.MapId)
		BattleFailed := 0
		if err != nil {
			if err.Error() == "Loading..." {
				monster, err = repositoryInstance.EnterMap(Config.PetId, Config.MapId)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		for {
			//
			StartTime := time.Now()
			result, isFinish := repositoryInstance.Fight(Config.SkillId, monster)
			if result == "11" {
				if isFinish {
					return nil
				} else {
					break
				}
			} else if result == "00" {
				plugin_log.Error("战斗失败，当前战斗失败 %d 次", BattleFailed)
				BattleFailed = BattleFailed + 1
				if BattleFailed > MaxBattleFailedTimes {
					plugin_log.Error("连续战斗失败 %d 次，当前战斗失败 %d 次", MaxBattleFailedTimes, BattleFailed)
					return errors.New(plugin_sdk_const.BATTLE_FAILED_TOO_MANY_TIMES)
				}
			} else if result == "10" {
				CostTime := time.Since(StartTime).Milliseconds()
				plugin_log.Debug("耗时 %d", CostTime)
				SleepTime := 3000 - CostTime
				time.Sleep(time.Duration(SleepTime) * time.Millisecond)
				continue
			} else {
				plugin_log.Error("战斗数据异常,重新进入地图")
				break
			}
		}
	}
}

func OpenDungeonInstance(Config *model.DungeonInstanceConfig) error {
	_, err := repositoryInstance.SpendSj(Config.MapId)
	if err != nil {
		if err.Error() == plugin_sdk_const.NOT_ENOUGH_SJ && Config.UseSj {
			plugin_log.Info("水晶不足,使用水晶卡进行开启")
			err = ArticleServiceInstance.UseSjk()
			if err != nil {
				return err
			}
			return OpenDungeonInstance(Config)
		} else {
			return err
		}
	}

	return err
}
