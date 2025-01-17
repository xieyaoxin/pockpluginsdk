package status

// FightStatus 挂机任务状态
var FightStatus = &fightStatus{
	BattleStatus:   NotReady,
	FbBattleStatus: NotReady,
	TtBattleStatus: NotReady,
}

const (
	NotReady     = "NOT_READY"
	Running      = "RUNNING"
	Waiting2Stop = "WAITING_TO_STOP"
)

type fightStatus struct {
	BattleStatus   string
	FbBattleStatus string
	TtBattleStatus string
}

func IsBattleRunning() bool {
	return !(FightStatus.BattleStatus == NotReady && FightStatus.FbBattleStatus == NotReady && FightStatus.TtBattleStatus == NotReady)
}

func IsBattleParsing() bool {
	return FightStatus.BattleStatus == Waiting2Stop || FightStatus.FbBattleStatus == Waiting2Stop || FightStatus.TtBattleStatus == Waiting2Stop
}

func IsBattleNotReady() bool {
	return FightStatus.BattleStatus == NotReady && FightStatus.FbBattleStatus == NotReady && FightStatus.TtBattleStatus == NotReady
}

func SetBattleStatus(status string) {
	FightStatus.BattleStatus = status
}

func SetTtBattleStatus(status string) {
	FightStatus.TtBattleStatus = status
}

func SetFbBattleStatus(status string) {
	FightStatus.FbBattleStatus = status
}
