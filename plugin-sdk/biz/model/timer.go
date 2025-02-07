package model

type TimerConfig struct {
	Enable                bool                       `json:"enable"`
	TimeTaskType          string                     `json:"type"`
	Schedule              int                        `json:"schedule"`
	Config                interface{}                `json:"config"`
	DropArticleConfig     DropArticleTimerConfig     `json:"drop_article_config"`
	DungeonInstanceConfig DungeonInstanceFightConfig `json:"dungeon_instance_config"`
}

type DropArticleTimerConfig struct {
	DropArticleList      []string `json:"drop_article_list"`
	DropArticleBlackList []string `json:"drop_article_black_list"`
}

type TimerHandleInterface interface {
	HandleTimer(TimerConfig)
}
