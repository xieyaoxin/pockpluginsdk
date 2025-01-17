package service

import (
	"pock_plugins/backend/service/impl"
)

// GetCarriedPets 查询身上的宠物信息
func (app *App) GetCarriedPetList() string {
	data, err := impl.PetServiceInstance.GetCarriedPetList()
	if err != nil {
		return buildFailedResponse(nil)
	}
	return buildSuccessResponse(data)
}

// GetPetDetail 获取宠物详情
func (app *App) GetPetDetail(PetId string) string {
	data, err := impl.PetServiceInstance.GetPetDetail(PetId)
	if err != nil {
		return buildFailedResponse(nil)
	}
	return buildSuccessResponse(data)
}

// GetPetSkillList 获取宠物技能列表
func (app *App) GetPetSkillList(PetId string) string {
	data, err := impl.PetServiceInstance.GetPetSkillList(PetId)
	if err != nil {
		return buildFailedResponse(nil)
	}
	return buildSuccessResponse(data)
}

// GetFarmedPets 查询牧场中的宠物
func (app *App) GetFarmedPets() string {
	data, err := impl.PetServiceInstance.GetFarmedPets()
	if err != nil {
		return buildFailedResponse(nil)
	}
	return buildSuccessResponse(data)
}

// SetBattlePet 设置主站宠物
func (app *App) SetBattlePet(PetId string) string {
	err := impl.PetServiceInstance.SetBattlePet(PetId)
	if err != nil {
		return buildFailedResponse(nil)
	}
	return buildSuccessResponse(nil)
}
