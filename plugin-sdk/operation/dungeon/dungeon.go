package tt

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/runner"

var DungeonServiceImplInstance = &dungeonServiceImpl{}

type dungeonServiceImpl struct {
}

func (inst *dungeonServiceImpl) StartTt() {
	err := runner.StartTask(runner.DUNGEON)
	if err != nil {
		return
	}
}
