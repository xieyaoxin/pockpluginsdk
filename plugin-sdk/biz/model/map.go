package model

type BattleMap struct {
	*PockBaseModel
	MapId          string
	MapName        string
	DifficultyFlag bool
}
