package timer

import plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"

type AutoTaskTimerHandler struct {
}

func (*AutoTaskTimerHandler) HandleTimer(OriginConfig interface{}) {
	plugin_sdk.TaskServiceInstance.StartAndFinishAllTask()
}
