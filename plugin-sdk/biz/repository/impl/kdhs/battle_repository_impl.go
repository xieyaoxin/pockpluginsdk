package kdhs

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"

	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"strconv"
	"strings"
	"time"
)

var BattleRepositoryImplInstance = &battleRepositoryImpl{}

type battleRepositoryImpl struct{}

func (battleRepositoryImpl) EnterMap(petId string) (*model.Monster, error) {
	return enterMap(petId)
}

func (battleRepositoryImpl) SelectAndEnterMap(mapId string, petId string) (*model.Monster, error) {
	params := util.InitParam()
	CallServerGetInterface("function/Pets_Mod.php", params)

	params["n"] = mapId
	params["type"] = "1"
	CallServerGetInterface("function/mapGate.php", params)

	params["p"] = petId
	params["mapid"] = mapId
	CallServerGetInterface("function/manymapgate.php", params)

	//result1 = CallServerGetInterface("function/Team_Mod.php", params)
	////params = util.InitParam()
	//params["type"] = "5"
	//result1 = CallServerGetInterface("function/mapGate.php", params)
	//log.Info("result is %s", result1)
	monster, err := enterMap(petId)
	if monster != nil {
		plugin_log.Info("当前怪物: %s 等级: %d, Hp: %d", monster.Name, monster.Level, monster.CurrentHp)
	}
	return monster, err
}

func (battleRepositoryImpl) FightOnce(SkillId string, monster *model.Monster) string {
	params := util.InitParam()
	params["id"] = SkillId
	params["g"] = monster.SkillId
	params["checkwg"] = "checked"
	params["rd"] = "0.13230130910911164"
	result := CallServerGetInterface("function/FightGate.php", params)
	resultArray := strings.Split(result, "#")
	if result == "" || len(resultArray) < 2 {
		plugin_log.Info("解析异常  重新进入战斗, 原始响应： %s", result)
		return result
	}
	// 计算怪物剩余血量
	leftHp, err := strconv.Atoi(strings.Split(resultArray[1], ",")[0])
	if err != nil {
		plugin_log.Info("解析怪物血量错误 原始信息:%s", result)
	}
	monster.CurrentHp = leftHp
	monster.CalculateCurrentHpRate()
	plugin_log.Info("当前怪物：%s. 剩余血量 %d, 百分比 %d %", monster.Name, monster.CurrentHp, monster.CurrentHpRate)
	// 判断
	if strings.Contains(resultArray[2], "受到了严重伤害，已经不能战斗") {
		return "00"
	}
	if strings.Contains(resultArray[2], "获得经验：") {
		plugin_log.Info(resultArray[2])
		return "11"
	}
	return "10"
}

func (instance *battleRepositoryImpl) CatchPet(monster *model.Monster, BallId string) bool {
	params := util.InitParam()
	params["pid"] = BallId
	catchResult := CallServerGetInterface("function/get.Catch.php", params)
	plugin_log.Info("遭遇宠物: %s  捕捉结果: %s", monster.Name, catchResult)
	return catchResult == "10"
}

//
//func (instance *BattleRepositoryImpl) CatchPet(mapId string, catchHpThreshold int, petId,skillId string, monster *model.Monster, BallId string) bool {
//	for monster.CurrentHpRate > catchHpThreshold {
//		result := instance.FightOnce(petId,skillId,monster )
//		// 战斗成功 / 战斗失败 -> 返回捕捉失败
//		if result == "00" || result == "11" {
//			return false
//		}
//	}
//	var catchResult string
//	if BallId != "" {
//		params := util.InitParam()
//		params["pid"] = ballId
//		catchResult = CallServerGetInterface("function/get.Catch.php", params)
//		log.Info("遭遇宠物: %s  捕捉结果: %s", monster.Name, catchResult)
//	} else {
//		log.Info("遭遇宠物: %s  找不到道具 %s", monsterName, ballName)
//		return false
//	}
//	// catchResult
//	if catchResult == "10" {
//		// 寄存
//		SaveUnBattle(property)
//	}
//	return catchResult == "10"
//}

// EnterMap 普通地图 进入战斗 获取战斗信息
func enterMap(petId string) (*model.Monster, error) {
	// 进入战斗
	params := util.InitParam()
	params["p"] = petId
	params["type"] = "1"
	result := CallServerGetInterface("function/Fight_Mod.php", params)
	//if strings.Contains(result, "继续31层，将收取300水晶，是否继续") {
	//	params["confirm31"] = "yes"
	//	result = CallServerGetInterface(&property, "function/Fight_Mod.php", params)
	//}
	var sleepTime int64
	if result == "不能获得宠物数据！" {

	}
	for strings.Contains(result, "loadtime") {
		time2Sleep := strings.Split(strings.Replace(strings.Replace(strings.Split(result, "loadtime")[3], "(", "/", 1), ")", "/", 1), "/")[1]
		sleepTime1, _ := strconv.ParseInt(time2Sleep, 10, 64)
		plugin_log.Info("等待 %s 秒后进入地图", time2Sleep)
		sleepTime = sleepTime1*1000 + 500
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		result = CallServerGetInterface("function/Fight_Mod.php", params)
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
		NatureType: monsterPropertyArray[2].(string), TotalHp: int(monsterPropertyArray[5].(float64)), CurrentHp: int(monsterPropertyArray[5].(float64)),
		CurrentHpRate: 100, SkillId: strconv.Itoa(int(monsterPropertyArray[11].(float64)))}, nil
}
