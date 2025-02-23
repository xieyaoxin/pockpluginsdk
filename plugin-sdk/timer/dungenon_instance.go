package timer

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
)

var DungeonStatus = status.NotReady

type DungeonInstanceTimerHandler struct {
}

func (*DungeonInstanceTimerHandler) HandleTimer(OriginConfig interface{}) {
	if DungeonStatus != status.NotReady {
		return
	}
	OriginBattleType := status.GetCurrentTaskType()
	if OriginBattleType == status.FUSION || OriginBattleType == status.NIRVANA || OriginBattleType == status.DUNGEON {
		plugin_log.Info("合神、涅槃、副本挂机时,跳过定时副本")
		return
	}
	OriginBattleStatus := status.GetBattleStatus()

	Config, err := Convert2DungeonInstanceFightConfig(OriginConfig)
	dungeonInstanceMapList := Config.MapList
	if dungeonInstanceMapList == nil || len(dungeonInstanceMapList) == 0 {
		return
	}
	if err != nil {
		plugin_log.Error("解析定时任务配置失败,%d", OriginConfig)
		return
	}

	// 检查副本状态
	DungeonOpenStatus := false
	for _, MapId := range dungeonInstanceMapList {
		open, _ := plugin_sdk.DungeonInstanceServiceImplInstance.GetDungeonStatus(MapId)
		if open {
			DungeonOpenStatus = true
			break
		}
	}
	if !DungeonOpenStatus {
		plugin_log.Info("副本未开启")
		return
	}
	if OriginBattleStatus == status.Running {
		// 暂停战斗
		status.SetBattleStatus(status.Parsing)
	} else if OriginBattleStatus == status.NotReady && DungeonStatus == status.NotReady {
		DungeonStatus = status.Running
		loopDungeon(Config, dungeonInstanceMapList)
		DungeonStatus = status.NotReady
		return
	}

	//
	select {
	case <-model.DungeonChannel:
		status.SetBattleStatus(OriginBattleStatus)
		DungeonStatus = status.Running
		loopDungeon(Config, dungeonInstanceMapList)
		model.FightChannel <- OriginBattleStatus
		DungeonStatus = status.NotReady
		break
	}
	//Articls, err1 := plugin_sdk.ArticleServiceInstance.QueryArticleList("")
	//if err1 != nil {
	//	plugin_log.Error("获取物品列表失败")
	//	return
	//}
	//if Config.DropArticleBlackList == nil || len(Config.DropArticleList) == 0 {
	//	return
	//}
	//if Config.DropArticleBlackList == nil {
	//	Config.DropArticleBlackList = []string{}
	//}
	//for _, Article := range Articls {
	//	DropFlag := util.SlicesLikeString(Config.DropArticleList, Article.Name)
	//	if DropFlag {
	//		if Article.ArticleType == "礼包类" || Article.ArticleType == "宝箱类" {
	//			plugin_log.Info("物品: %s 为礼包类 跳过丢弃", Article.Name)
	//			continue
	//		}
	//		BlackFlag := util.SlicesLikeString(Config.DropArticleBlackList, Article.Name)
	//		if !BlackFlag {
	//			plugin_log.Info("丢弃物品: %s,物品Id: %s,物品数量: %d", Article.Name, Article.ID, Article.ArticleCount)
	//		}
	//	}
	//}
	//plugin_sdk.DungeonInstanceServiceImplInstance.FightDungeonOnce(Config)

}

func Convert2DungeonInstanceFightConfig(OriginConfig interface{}) (*model.DungeonInstanceTimerConfig, error) {
	Config := &model.DungeonInstanceTimerConfig{}
	input, err := json.Marshal(OriginConfig)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(input, Config)
	if err != nil {
		return nil, err
	}
	return Config, nil
}

func loopDungeon(Config *model.DungeonInstanceTimerConfig, dungeonInstanceMapList []string) {
	PetId, SkillId := plugin_sdk.InitBattlePet(Config.PetId, Config.PetName, Config.SkillId, Config.SkillName)
	if PetId == "" || SkillId == "" {
		logrus.Info("未配置挂机宠物和技能")
		return
	}
	for _, MapId := range dungeonInstanceMapList {
		DungeonConfig := &model.DungeonInstanceFightConfig{
			PetId:      PetId,
			SkillId:    SkillId,
			MapId:      MapId,
			ForceFight: false,
			UseSj:      false,
		}
		err2 := plugin_sdk.DungeonInstanceServiceImplInstance.FightDungeonOnce(DungeonConfig)
		if err2 != nil {
			continue
		}
	}
}
