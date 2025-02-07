package chains

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	model2 "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain/chain_model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain/config"
)

func FusionChain() {
	currentUser := status.GetLoginUser()
	FusionChainConfig := config.GetFusionConfig(currentUser.LoginName)
	MergeCount := 0
	MergeConfig := chain_model.CopyFusionConfig(*FusionChainConfig)
	// 初始化缓存
	plugin_sdk.InitMergeArticleCache()
	for {
		GodList := []*model2.Pet{}

		petsInFarm, err := plugin_sdk.PetServiceInstance.GetFarmPets()
		petsInBody, err := plugin_sdk.PetServiceInstance.GetCarriedPetList()
		pets := append(petsInFarm, petsInBody...)
		if err != nil {
			return
		}
		for _, Pet := range pets {
			if Pet.Name == "小神龙琅玡" {
				GodList = append(GodList, Pet)
			}
		}
		GodCount := len(GodList)
		plugin_log.Info("当前牧场中小神数量 %d", GodCount)
		if FusionChainConfig.Finish.GodCount > 0 && GodCount >= FusionChainConfig.Finish.GodCount {
			plugin_log.Info("达到预期数量 %d,合神结束", FusionChainConfig.Finish.GodCount)
			break
		}
		_, err = plugin_sdk.MergeGod(*MergeConfig)
		if err != nil {
			plugin_log.Error("合神失败 原因: %s", err.Error())
			return
		}
		MergeCount = MergeCount + 1
		plugin_log.Info("当前合成数量 %d", MergeCount)
		if MergeCount >= FusionChainConfig.Finish.MergeCount {
			plugin_log.Info("达到预期数量 %d,合神结束", FusionChainConfig.Finish.MergeCount)
			break
		}
	}
}
