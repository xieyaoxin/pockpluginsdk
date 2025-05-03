package fusion

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/runner"

var FusionImplInstance = &fusionImpl{}

type fusionImpl struct {
}

func (inst *fusionImpl) Start() {
	err := runner.StartTask(runner.MERGE)
	if err != nil {
		return
	}
}
