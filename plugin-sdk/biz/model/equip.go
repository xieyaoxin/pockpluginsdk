package model

type Equip struct {
	*PockBaseModel
	EquipId    string
	Name       string
	Position   string
	Strengthen int64
	Effect     Effect
}

type Effect struct {
	EffectType string
	Figure     int64
}
