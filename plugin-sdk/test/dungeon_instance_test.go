package test

import (
	"fmt"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/chain"
	"testing"
)

func TestDungeonInstanceFightOnce(t *testing.T) {
	GetLoginUser()
	Config := &model.DungeonInstanceFightConfig{
		PetId:      "8512800",
		SkillId:    "799",
		MapId:      "144",
		ForceFight: true,
		UseSj:      false,
	}

	err := plugin_sdk.DungeonInstanceServiceImplInstance.FightDungeonOnce(Config)
	if err != nil {
		fmt.Printf("error is %s", err.Error())
		return
	}
}
