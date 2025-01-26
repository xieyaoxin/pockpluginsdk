package repository

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"

type EquipRepository interface {
	// GetEquipList 获取装备
	GetEquipList(PetId string) []*model.Equip
	// StrengthEquip 强化装备
	StrengthEquip(EquipId string, EquipPid string, DragonBallId string) bool
	// OffEquip 脱装备
	OffEquip(PetId string, EquipPid string) bool

	GetEquip(EquipId string) *model.Equip
}
