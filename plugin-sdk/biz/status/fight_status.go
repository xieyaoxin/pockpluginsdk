package status

// FightStatus 挂机任务状态

const (
	NotReady     = "NOT_READY"
	Running      = "RUNNING"
	Parsing      = "PARSING"
	Waiting2Stop = "WAITING_TO_STOP"
)

var battleStatus = "NOT_READY"

func IsBattleRunning() bool {
	return battleStatus == Running
}

func IsParsing() bool {
	return battleStatus == Waiting2Stop || battleStatus == Parsing
}

func IsBattleNotReady() bool {
	return battleStatus == NotReady
}

func SetBattleStatus(status string) {
	battleStatus = status
}

func GetBattleStatus() string {
	return battleStatus
}
