package model

type TtConfig struct {
	PetId      string `json:"pet_id,omitempty"`
	PetName    string `json:"pet_name,omitempty"`
	SkillId    string `json:"skill_id,omitempty"`
	SkillName  string `json:"skill_name,omitempty"`
	DropEquip  bool   `json:"drop_equip"`
	DropZll    bool   `json:"drop_zll"`
	AutoUseSjk bool   `json:"auto_use_sjk"`
	MaxLevel   int    `json:"max_level"`
	LoopTt     bool   `json:"loop_tt"`
}
