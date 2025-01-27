package model

type BattleConfig struct {
	PetId              string   `json:"pet_id,omitempty"`
	SkillId            string   `json:"skill_id,omitempty"`
	MapId              string   `json:"map_id,omitempty"`
	Difficulty         string   `json:"difficulty,omitempty"`
	SkipMonsters       []string `json:"skip_monsters,omitempty"`
	CatchPets          []string `json:"catch_pets,omitempty"`
	RunWhenCatchFailed bool     `json:"run_when_catch_failed,omitempty"`
	RunWhenNotCatch    bool     `json:"run_when_not_catch,omitempty"`
	Balls              []string `json:"balls,omitempty"`
	Rubbish            []string `json:"rubbish,omitempty"`
	CatchHpThreshold   int      `json:"catch_hp_threshold,omitempty"`
	SaveAfterCatch     bool     `json:"save_after_catch"`
}
