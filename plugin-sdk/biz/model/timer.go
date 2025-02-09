package model

var DungeonChannel = make(chan string)
var FightChannel = make(chan string)

type TimerConfig struct {
	Enable       bool                 `json:"enable"`
	TimeTaskType string               `json:"type"`
	Schedule     int64                `json:"schedule"`
	Config       interface{}          `json:"config"`
	Handle       TimerHandleInterface `json:"-"`
}

type TimerHandleInterface interface {
	HandleTimer(interface{})
}

type DropArticleTimerConfig struct {
	DropArticleList      []string `json:"drop_article_list"`
	DropArticleBlackList []string `json:"drop_article_black_list"`
}

type DungeonInstanceTimerConfig struct {
	PetId     string   `json:"pet_id"`
	PetName   string   `json:"pet_name"`
	SkillId   string   `json:"skill_id"`
	SkillName string   `json:"skill_name"`
	MapList   []string `json:"map_list"`
}
