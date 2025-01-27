package plugin_sdk_const

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
)

var monsterBallMap = make(map[string][]string)

func initMonsterBallMap() {
	monsterBallMap["涅磐兽"] = []string{"涅槃兽·精灵球(绑定)", "涅盘兽·精灵球", "涅盘兽·精灵球（新手版）",
		"涅盘兽·大师精灵球（绑定）", "涅盘兽·大师精灵球"}
	monsterBallMap["涅槃兽（亥）"] = []string{"涅盘兽·精灵球", "涅盘兽地图·大师精灵球（绑定）"}
}

func GetBallByMonster(monsterName string, balls []string) []string {
	BallNames := []string{}
	if BallName, exists := monsterBallMap[monsterName]; exists {
		if balls != nil && len(balls) > 0 {
			for _, ConfigBall := range balls {
				if utils.SlicesContainsString(BallName, ConfigBall) {
					BallNames = append(BallNames, ConfigBall)
				}
			}
		} else {
			BallNames = BallName
		}
	} else {
		BallNames = append(BallNames, monsterName+"·精灵球")
	}
	return BallNames
}
