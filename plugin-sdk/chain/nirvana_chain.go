package chains

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain/chain_model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain/config"
)

func NirvanaChain() {
	currentUser := status.GetLoginUser()
	GetNirvanaConfig := config.GetNirvanaConfig(currentUser.LoginName)
	NirvanaCount := 0
	NirvanaConfig := chain_model.CopyNirvanaConfig(*GetNirvanaConfig)
	for {
		if NirvanaCount >= GetNirvanaConfig.Finish.NirvanaCount && GetNirvanaConfig.Finish.NirvanaCount > 0 {
			plugin_log.Info("已涅槃 %d 次，涅槃数量达到预期 %d", NirvanaCount, GetNirvanaConfig.Finish.NirvanaCount)
			break
		}
		_, err := plugin_sdk.NirvanaServiceImplInstance.Nirvana(NirvanaConfig)
		if err != nil {
			return
		}
		NirvanaCount = NirvanaCount + 1
		Pet := plugin_sdk.PetServiceInstance.GetBattlePet()
		plugin_log.Info("第 %d 次。涅槃成功 当前宠物 %s 成长: %f", NirvanaCount, Pet.Name, Pet.Cc)
	}
}
