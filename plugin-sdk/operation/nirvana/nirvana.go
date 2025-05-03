package nirvana

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/runner"

var NirvanaServiceImplInstance = &NirvanaServiceImpl{}

type NirvanaServiceImpl struct {
}

func (inst *NirvanaServiceImpl) Start() {
	err := runner.StartTask(runner.NIRVANA)
	if err != nil {
		return
	}
}
