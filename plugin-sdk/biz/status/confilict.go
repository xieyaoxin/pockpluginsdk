package status

const (
	// 无
	NONE = iota
	// 合神
	FUSION
	// 涅槃
	NIRVANA
	// 挂机
	BATTLE
	// 副本
	DUNGEON
	// 通天
	TT
)

func GetConflictTask() bool {
	return IsBattleRunning() || IsParsing()
}

var currentTask = NONE

func SetTaskType(TaskType int) {
	currentTask = TaskType
}
