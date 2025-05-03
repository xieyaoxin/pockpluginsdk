package chain_model

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
)

type FusionChainConfig struct {
	model.MergeGodConfig
	Finish FusionFinishConfig `json:"finish"`
}

type FusionFinishConfig struct {
	MergeCount int `json:"merge_count"`
	GodCount   int `json:"god_count"`
}

func CopyFusionConfig(config FusionChainConfig) *model.MergeGodConfig {

	return &model.MergeGodConfig{
		MainPet:   config.MainPet,
		AteDragon: config.AteDragon,
		EatDragon: config.EatDragon,
	}
}
