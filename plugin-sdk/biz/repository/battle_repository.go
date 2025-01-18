package repository

import (
	"github.com/xieyaoxin/plugin-sdk/plugin-sdk/biz/model"
)

type BattleRepository interface {
	// SelectAndEnterMap 进入地图 遭遇怪物
	SelectAndEnterMap(mapId string, petId string) (*model.Monster, error)
	// FightOnce 执行一次攻击.返回状态: 11: 战斗成功-战斗结束; 10 战斗成功-战斗未结束; 00 战斗失败
	FightOnce(SkillId string, monster *model.Monster) string
	// CatchPet 执行一次捕捉
	CatchPet(monster *model.Monster, BallId string) bool
}
