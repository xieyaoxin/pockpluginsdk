package model

type BattleConfig struct {
	PetId              string
	SkillId            string
	MapId              string
	Difficulty         string
	SkipMonsters       []string
	CatchPets          []string
	RunWhenCatchFailed bool
	RunWhenNotCatch    bool
	Balls              []string
	Rubbish            []string
	CatchHpThreshold   int
}
