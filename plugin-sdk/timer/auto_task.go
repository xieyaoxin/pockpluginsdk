package timer

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/task"
)

type AutoTaskTimerHandler struct {
}

func (*AutoTaskTimerHandler) HandleTimer(OriginConfig interface{}) {
	plugin_sdk.TaskServiceInstance.StartAndFinishAllTask()
}
