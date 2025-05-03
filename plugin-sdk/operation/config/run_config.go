package config

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/chain"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/runner"
)

func init() {
	//InitMergeArticleCache()
	runner.HandlerMap[runner.TT_RUNNING] = chain.TTStateMachineInstance
	runner.HandlerMap[runner.BATTLE_RUNNING] = BattleStateMachineInstance

}
