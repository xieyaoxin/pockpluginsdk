package chain_model

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
)

type NirvanaChainConfig struct {
	model.NirvanaConfig
	Finish NirvanaFinishConfig `json:"finish"`
}

type NirvanaFinishConfig struct {
	NirvanaCount int `json:"nirvana_count"`
}

func CopyNirvanaConfig(config NirvanaChainConfig) *model.NirvanaConfig {
	return &model.NirvanaConfig{
		MainPet:      config.MainPet,
		AtePet:       config.AtePet,
		NirvanaPet:   config.NirvanaPet,
		ProtectType1: config.ProtectType1,
		ProtectType2: config.ProtectType2,
	}
}
