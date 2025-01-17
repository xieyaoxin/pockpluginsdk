package model

type Monster struct {
	*PockBaseModel
	Name          string
	TotalHp       int
	CurrentHp     int
	NatureType    string
	SkillId       string
	CurrentHpRate int
	Level         int
}

func (monster *Monster) CalculateCurrentHpRate() {
	if monster.TotalHp == 0 {
		monster.CurrentHpRate = 1
		return
	}
	rate := monster.CurrentHpRate * 100 / monster.TotalHp
	monster.CurrentHpRate = rate
}
