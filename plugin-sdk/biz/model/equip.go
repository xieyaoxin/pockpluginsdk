package model

type Equip struct {
	*PockBaseModel
	EquipId    string
	EquipPId   string
	Name       string
	Position   string
	Strengthen int64
	Effect     Effect
}

type Effect struct {
	EffectType string
	Figure     int64
}
