package repository

type TtRepository interface {
	// 进地图
	EnterTt() string
	ShouldPaySj(action string) string
	Pay30SJ(string) bool
}
