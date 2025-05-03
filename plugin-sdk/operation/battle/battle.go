package battle

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/runner"

var TtServiceImplInstance = &ttServiceImpl{}

type ttServiceImpl struct {
}

func (inst *ttServiceImpl) StartTt() {
	err := runner.StartTask(runner.BATTLE)
	if err != nil {
		return
	}
}
