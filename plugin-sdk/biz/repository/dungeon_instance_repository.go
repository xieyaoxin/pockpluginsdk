package repository

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"

type DungeonInstanceRepository interface {
	// GetDungeonInstanceStatus 获取当前副本状态 返回值：当前关卡,倒计时
	GetDungeonInstanceStatus(MapId string) (int, int)
	// SpendSj 消耗进入副本的水晶
	SpendSj(MapId string) (bool, error)

	// EnterMap 进入副本,进入副本失败时 返回对应的异常信息
	EnterMap(PetId, MapId string) (*model.Monster, error)

	// FightOnce 执行一次攻击.返回状态: 11: 战斗成功-战斗结束; 10 战斗成功-战斗未结束; 00 战斗失败 ; true: 副本结束 false: 副本未结束
	Fight(SkillId string, monster *model.Monster) (string, bool)
}
