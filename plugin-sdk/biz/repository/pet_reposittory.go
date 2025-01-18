package repository

import (
	"github.com/xieyaoxin/plugin-sdk/plugin-sdk/biz/model"
)

type PetRepository interface {
	// GetCarriedPets 查询身上的宠物信息
	GetCarriedPets() ([]*model.Pet, error)
	// GetPetDetail 获取宠物详情
	GetPetDetail(PetId string) (*model.Pet, error)
	// GetPetSkillList 获取宠物技能列表
	GetPetSkillList(PetId string) ([]*model.Skill, error)
	// GetFarmedPets 查询牧场中的宠物
	GetFarmedPets() ([]*model.Pet, error)
	// SetBattlePet 设置主站宠物
	SetBattlePet(PetId string) error
	// CarryPet 携带宠物到身上
	CarryPet(PetId string) error
	// SavePet 存储宠物到牧场
	SavePet(PetId string) error
}
