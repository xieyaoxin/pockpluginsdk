package chains

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain/config"
)

var DungeonInstanceReportCallbackInstance = &DungeonInstanceReportCallback{}

type DungeonInstanceReportCallback struct {
}

func (*DungeonInstanceReportCallback) Callback(input interface{}) {
}

func (*DungeonInstanceReportCallback) StopCallback() {

}
func (*DungeonInstanceReportCallback) OverWriteDungeon() {

}

func DungeonInstanceChain() {
	currentUser := status.GetLoginUser()
	DungeonInstanceConfig := config.GetDungeonInstanceConfig(currentUser.LoginName)
	PetId, SkillId := plugin_sdk.InitBattlePet(DungeonInstanceConfig.PetId, DungeonInstanceConfig.PetName, DungeonInstanceConfig.SkillId, DungeonInstanceConfig.SkillName)
	battleConfig := &model.DungeonInstanceConfig{
		DungeonInstanceFightConfig: model.DungeonInstanceFightConfig{
			PetId:      PetId,
			SkillId:    SkillId,
			MapId:      DungeonInstanceConfig.MapId,
			ForceFight: DungeonInstanceConfig.ForceFight,
			UseSj:      DungeonInstanceConfig.UseSj,
		},
		FightTimes: DungeonInstanceConfig.FightTimes,
	}
	plugin_sdk.DungeonInstanceServiceImplInstance.FightDungeon(battleConfig, DungeonInstanceReportCallbackInstance)
}
