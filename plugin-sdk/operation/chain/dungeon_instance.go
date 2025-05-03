package chain

import (
	"context"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/battle"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/chain/spi/config"
)

var DungeonStateMachineInstance = &DungeonStateMachineApi{}

type DungeonStateMachineApi struct{}

var dungeonCtx context.Context
var dungeonCancel context.CancelFunc

func (DungeonStateMachineApi) HandleFinishEvent() {
	dungeonCancel()
}

func (DungeonStateMachineApi) HandleStartEvent() {
	dungeonCtx, dungeonCancel = context.WithCancel(context.Background())
	go startDungeonTask(dungeonCtx)
}

func startDungeonTask(context context.Context) {
	//dungeonConfig := initDungeonConfig()
}

func getDungeonTask() *model.DungeonInstanceConfig {
	currentUser := status.GetLoginUser()
	DungeonInstanceConfig := config.GetDungeonInstanceConfig(currentUser.LoginName)
	PetId, SkillId := battle.InitBattlePet(DungeonInstanceConfig.PetId, DungeonInstanceConfig.PetName, DungeonInstanceConfig.SkillId, DungeonInstanceConfig.SkillName)
	return &model.DungeonInstanceConfig{
		DungeonInstanceFightConfig: model.DungeonInstanceFightConfig{
			PetId:      PetId,
			SkillId:    SkillId,
			MapId:      DungeonInstanceConfig.MapId,
			ForceFight: DungeonInstanceConfig.ForceFight,
			UseSj:      DungeonInstanceConfig.UseSj,
		},
		FightTimes: DungeonInstanceConfig.FightTimes,
	}

}
