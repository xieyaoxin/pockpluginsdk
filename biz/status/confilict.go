package status

func GetConflictTask() bool {
	return IsBattleRunning() || IsBattleParsing()
}
