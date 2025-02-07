package chain_model

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
)

type DungeonChainConfig struct {
	model.DungeonInstanceConfig
	PetName   string `json:"pet_name"`
	SkillName string `json:"skill_name"`
}
