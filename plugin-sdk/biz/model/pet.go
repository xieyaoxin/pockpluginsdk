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

func PetSliceRemoveItem(List []*Pet, Item *Pet) []*Pet {
	TempList := []*Pet{}
	for _, Item1 := range List {
		if Item != Item1 {
			TempList = append(TempList, Item1)
		}
	}
	return TempList
}
