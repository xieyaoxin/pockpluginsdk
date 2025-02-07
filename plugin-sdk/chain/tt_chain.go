package chains

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain/config"
)

var TtReportCallbackInstance = &TtReportCallback{}

type TtReportCallback struct {
}

func (*TtReportCallback) Callback(input interface{}) {
}

func (*TtReportCallback) StopCallback() {

}
func (*TtReportCallback) OverWriteTt() {

}

func TtChain() {
	currentUser := status.GetLoginUser()
	ttConfig := config.GetTtConfig(currentUser.LoginName)
	PetId, SkillId := plugin_sdk.InitBattlePet(ttConfig.PetId, ttConfig.PetName, ttConfig.SkillId, ttConfig.SkillName)
	battleConfig := &model.TtConfig{
		PetId:      PetId,
		SkillId:    SkillId,
		DropEquip:  ttConfig.DropEquip,
		DropZll:    ttConfig.DropZll,
		AutoUseSjk: ttConfig.AutoUseSjk,
		MaxLevel:   ttConfig.MaxLevel,
		LoopTt:     ttConfig.LoopTt,
	}
	plugin_sdk.TtServiceImplInstance.StartTt(battleConfig, TtReportCallbackInstance)
}
