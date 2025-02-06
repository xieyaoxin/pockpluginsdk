package model

type DungeonInstanceFightConfig struct {
	PetId      string `json:"pet_id"`
	SkillId    string `json:"skill_id"`
	MapId      string `json:"map_id"`
	ForceFight bool   `json:"force_fight"`
	UseSj      bool   `json:"use_sj"`
}

type DungeonInstanceConfig struct {
	DungeonInstanceFightConfig
	FightTimes int `json:"fight_times"`
}
