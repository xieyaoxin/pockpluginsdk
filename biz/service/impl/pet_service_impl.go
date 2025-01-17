package impl

import (
	"pock_plugins/backend/model"
	"pock_plugins/backend/repository"
)

var PetServiceInstance = &petService{}

type petService struct {
}

// GetCarriedPets 查询身上的宠物信息
func (*petService) GetCarriedPetList() ([]*model.Pet, error) {
	return repository.GetPetRepository().GetCarriedPets()
}

// GetPetDetail 获取宠物详情
func (*petService) GetPetDetail(PetId string) (*model.Pet, error) {
	return repository.GetPetRepository().GetPetDetail(PetId)
}

// GetPetSkillList 获取宠物技能列表
func (*petService) GetPetSkillList(PetId string) ([]*model.Skill, error) {
	return repository.GetPetRepository().GetPetSkillList(PetId)
}

// GetFarmedPets 查询牧场中的宠物
func (*petService) GetFarmedPets() ([]*model.Pet, error) {
	return repository.GetPetRepository().GetFarmedPets()
}

// SetBattlePet 设置主站宠物
func (*petService) SetBattlePet(PetId string) error {
	return repository.GetPetRepository().SetBattlePet(PetId)
}

func (*petService) SaveUnBattlePet() error {
	list, _ := PetServiceInstance.GetCarriedPetList()
	for _, pet := range list {
		if !pet.IsBattle {
			err := repository.GetPetRepository().SavePet(pet.Id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
