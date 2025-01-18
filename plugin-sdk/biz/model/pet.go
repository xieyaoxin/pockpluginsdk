package model

type Pet struct {
	*PockBaseModel
	Id       string
	Name     string
	Cc       float64
	Level    int64
	IsBattle bool
}

type Skill struct {
	*PockBaseModel
	SkillName  string
	SkillLevel string
	SkillId    string
}
