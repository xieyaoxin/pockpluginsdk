package kdhs

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"strings"
)

var TtRepositoryImplKdhsInstance = &ttRepositoryImpl{}

type ttRepositoryImpl struct{}

func (*ttRepositoryImpl) EnterTt() string {
	params := util.InitParam()
	params["n"] = "126"
	map1 := CallServerGetInterface("function/Team_Mod.php", params)
	ttArray := strings.Split(map1, "\n")
	CurrentFloor := ""
	for lineNumber := range ttArray {
		line := ttArray[lineNumber]
		if strings.Contains(line, "当前关卡") && !strings.Contains(line, "#tgt#") && !strings.Contains(line, "玩家每天有一次进入通天塔") {
			CurrentFloor = strings.Split(strings.Split(line, "<")[1], ">")[1]
			plugin_log.Info(CurrentFloor)
			break
		}
	}
	return CurrentFloor
}

func (*ttRepositoryImpl) ShouldPaySj(action string) string {
	params := util.InitParam()
	params["op"] = "tgfight"
	if action != "" {
		params["action"] = "do"
	}
	result := CallServerGetInterface("function/ttGate.php", params)
	return result
}

func (*ttRepositoryImpl) Pay30SJ(PetId string) bool {
	params := util.InitParam()
	params["p"] = PetId
	params["type"] = "1"
	params["confirm31"] = "yes"
	// 进入地图
	r1 := CallServerGetInterface("function/Fight_Mod.php", params)
	return !strings.Contains(r1, "水晶不够，扣取失败")
}
