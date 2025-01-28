package test

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"testing"
)

//8206429
//

func TestTtFight(t *testing.T) {
	GetLoginUser()
	TtConfig := &model.TtConfig{
		PetId:   "8206429",
		SkillId: "852",
		LoopTt:  true,
	}
	plugin_sdk.TtServiceImplInstance.StartTtLoop(TtConfig, nil)
	HoldOn()
}
