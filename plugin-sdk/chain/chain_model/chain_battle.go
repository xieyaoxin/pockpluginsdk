package chain_model

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"

type BattleChainConfig struct {
	model.BattleConfig
	SkillName    string              `json:"skill_name,omitempty"`
	PetName      string              `json:"pet_name,omitempty"`
	FinishConfig *BattleFinishConfig `json:"finish_config"`
}

type BattleFinishConfig struct {
	// 战斗次数 <= 0 时跳过判断
	BattleCount int `json:"battle_count"`
	// 捕捉宠物数量 <= 0 时候跳过判断
	CatchPet int `json:"catch_pet"`
}
