package chain_model

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
)

type TtChainConfig struct {
	model.TtConfig
	SkillName string `json:"skill_name,omitempty"`
	PetName   string `json:"pet_name,omitempty"`
}
