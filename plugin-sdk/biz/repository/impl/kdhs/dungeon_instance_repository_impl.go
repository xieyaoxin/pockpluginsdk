package kdhs

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_sdk_const"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"strconv"
	"strings"
	"time"
)

var DungeonInstanceRepositoryKdhsImplInstance = &dungeonInstanceRepositoryKdhsImpl{}

type dungeonInstanceRepositoryKdhsImpl struct{}

func (*dungeonInstanceRepositoryKdhsImpl) GetDungeonInstanceStatus(MapId string) (int, int) {
	//http://43.248.129.148:567/function/fb_Mod.php?mapid=127
	Params := InitParam()
	Params["mapid"] = MapId
	result := CallServerGetInterface("/function/fb_Mod.php", Params)
	CurrentStage := -1
	CountDown := -1
	lines := strings.Split(result, "\r")
	for _, line := range lines {
		// 获取副本倒计时
		if strings.Contains(line, "副本开启倒计时") {
			//	<td height="20">副本开启倒计时：已开启</td>
			countDownsMessages := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "：", "<"), "：", "<"), "<")
			CountDownStr := countDownsMessages[8]
			if CountDownStr == "已开启" {
				CountDown = 0
			} else {
				CountDownStr = strings.ReplaceAll(CountDownStr, "秒", "")
				CountDown1, err := strconv.Atoi(CountDownStr)
				if err != nil {
					return CurrentStage, CountDown
				}
				CountDown = CountDown1
			}

		}
		// 获取当前关卡
		if strings.Contains(line, "当前进度") {
			//	<td height="20">副本开启倒计时：已开启</td>
			Messages := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "：", "<"), "：", "<"), "<")
			CurrentLevelStr := Messages[18]
			CurrentStage1, err := strconv.Atoi(CurrentLevelStr)
			if err != nil {
				return CurrentStage, CountDown
			}
			CurrentStage = CurrentStage1
		}
	}
	return CurrentStage, CountDown
}

func (*dungeonInstanceRepositoryKdhsImpl) EnterMap(PetId, MapId string) (*model.Monster, error) {
	// 获取地图信息
	params1 := util.InitParam()
	params1["mapid"] = MapId
	params1["p"] = PetId
	result1 := CallServerGetInterface("function/getmap.php", params1)
	if result1 != "10" {
		plugin_log.Error("获取地图信息失败")
		return nil, errors.New(plugin_sdk_const.GET_DUNGEON_MAP_FAILED)
	}
	// 进入战斗
	params := util.InitParam()

	params["p"] = PetId
	params["mapid"] = MapId
	result := CallServerGetInterface("function/fbfight_Mod.php", params)
	var sleepTime int64
	if result == "不能获得宠物数据！" {

	}
	for strings.Contains(result, "loadtime") {
		time2Sleep := strings.Split(strings.Replace(strings.Replace(strings.Split(result, "loadtime")[3], "(", "/", 1), ")", "/", 1), "/")[1]
		sleepTime1, _ := strconv.ParseInt(time2Sleep, 10, 64)
		plugin_log.Info("等待 %s 秒后进入地图", time2Sleep)
		sleepTime = sleepTime1*1000 + 500
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		result = CallServerGetInterface("function/fbfight_Mod.php", params)
	}
	// 获取怪物属性
	lines := strings.Split(result, "\n")
	var monsterPropertyArray []any
	for lineNumber := range lines {
		line := lines[lineNumber]
		if strings.Contains(line, "gg=") {
			monsterMessage := strings.Split(strings.Split(line, "=")[1], ";")[0]
			monsterMessage = strings.Replace(monsterMessage, "'", "\"", -1)
			monsterPropertyArray = util.String2JsonArray(monsterMessage)
		}
	}
	if len(monsterPropertyArray) < 12 {
		plugin_log.Error("进入地图失败 原因是: %s", result)
		return &model.Monster{}, errors.New(result)
	}
	return &model.Monster{Name: monsterPropertyArray[0].(string), Level: int(monsterPropertyArray[1].(float64)),
		NatureType: monsterPropertyArray[2].(string), TotalHp: int64(monsterPropertyArray[5].(float64)), CurrentHp: int64(monsterPropertyArray[5].(float64)),
		CurrentHpRate: 100, SkillId: strconv.Itoa(int(monsterPropertyArray[11].(float64)))}, nil
}

func (*dungeonInstanceRepositoryKdhsImpl) SpendSj(MapId string) (bool, error) {
	//http://43.248.129.148:567/function/mapGate.php?type=4&n=144
	Params := InitParam()
	Params["type"] = "4"
	Params["n"] = MapId
	result := CallServerGetInterface("/function/mapGate.php", Params)
	switch result {
	case "11":
		return true, nil
	case "3":
		return false, errors.New(plugin_sdk_const.NOT_ENOUGH_SJ)
	default:
		return false, errors.New(plugin_sdk_const.UNKNOWN_ERROR)
	}

}

// FightOnce 执行一次攻击.返回状态: 11: 战斗成功-战斗结束; 10 战斗成功-战斗未结束; 00 战斗失败 ; true: 副本结束 false: 副本未结束
func (*dungeonInstanceRepositoryKdhsImpl) Fight(SkillId string, monster *model.Monster) (string, bool) {
	params := util.InitParam()
	params["id"] = SkillId
	params["g"] = monster.SkillId
	params["checkwg"] = "checked"
	params["rd"] = "0.13230130910911164"
	result := CallServerGetInterface("function/fbfightGate.php", params)
	resultArray := strings.Split(result, "#")
	if result == "" || len(resultArray) < 2 {
		plugin_log.Info("解析异常  重新进入战斗, 原始响应： %s", result)
		return result, false
	}
	// 计算怪物剩余血量
	leftHpStr := strings.Split(resultArray[1], ",")[0]
	//var leftHp int64
	//var err error
	if strings.Contains(leftHpStr, "E+") {
		leftHpStr = strings.ReplaceAll(leftHpStr, "E+", "e")
		//leftHp,err := strconv.ParseInt(leftHpStr,)
	}
	leftHpFloat, err := strconv.ParseFloat(leftHpStr, 64)
	leftHp := int64(leftHpFloat)
	if err != nil {
		plugin_log.Info("解析怪物血量错误 原始信息:%s", result)
	}
	monster.CurrentHp = leftHp
	monster.CalculateCurrentHpRate()
	plugin_log.Info("当前怪物：%s. 剩余血量 %d, 百分比 %d  %%%%", monster.Name, monster.CurrentHp, monster.CurrentHpRate)
	// 判断
	if strings.Contains(resultArray[2], "受到了严重伤害，已经不能战斗") {
		return "00", false
	}
	if strings.Contains(resultArray[2], "获得经验：") {
		plugin_log.Info(resultArray[2])
		isFinish := strings.Contains(result, "#end#")
		return "11", isFinish
	}
	return "10", false
}
