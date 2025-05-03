package tt

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/runner"

var NirvanaServiceImplInstance = &NirvanaServiceImpl{}

type NirvanaServiceImpl struct {
}

func (inst *NirvanaServiceImpl) StartTt() {
	err := runner.StartTask(runner.NIRVANA)
	if err != nil {
		return
	}
}
