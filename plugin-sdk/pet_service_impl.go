package plugin_sdk

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"strings"
)

var PetServiceInstance = &petService{}

type petService struct {
}

// GetCarriedPets 查询身上的宠物信息
func (*petService) GetBattlePet() *model.Pet {
	pets, _ := repository.GetPetRepository().GetCarriedPets()
	for _, pet := range pets {
		if pet.IsBattle {
			return pet
		}
	}
	return nil
}

// GetCarriedPets 查询身上的宠物信息
func (*petService) GetCarriedPetList() ([]*model.Pet, error) {
	Pets, err := repository.GetPetRepository().GetCarriedPets()
	if Pets != nil {
		for _, Pet := range Pets {
			Pet.Carried = true
		}
	}
	return Pets, err
}

// GetPetDetail 获取宠物详情
func (*petService) GetPetDetail(PetId string) (*model.Pet, error) {
	return repository.GetPetRepository().GetPetDetail(PetId)
}

// GetPetSkillList 获取宠物技能列表
func (*petService) GetPetSkillList(PetId string) ([]*model.Skill, error) {
	return repository.GetPetRepository().GetPetSkillList(PetId)
}

// SetBattlePet 设置主站宠物
func (inst *petService) SetBattlePet(PetId string) error {
	result, err := repository.GetPetRepository().SetBattlePet(PetId)
	if result {
		return nil
	}
	// 宠物在牧场中
	if result == false && err == nil {
		err2 := inst.CarryPetInFarm(PetId)
		if err2 != nil {
			return err2
		}
		result, err = repository.GetPetRepository().SetBattlePet(PetId)
		return err
	}
	return err
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

func (*petService) GetFarmPets() ([]*model.Pet, error) {
	return repository.GetPetRepository().GetFarmedPets()
}

func (*petService) CarryPetInFarm(PetId string) error {
	return repository.GetPetRepository().CarryPet(PetId)
}

func (*petService) GetPetEvaluateArticle(Pet *model.Pet, Route string) *model.Article {
	if RouteArticleMap, exists := PetEvaluateRouteArticleMap[Pet.Name]; exists {
		return RouteArticleMap[Route]
	} else {
		RouteArticle1, RouteArticle2 := repository.GetFusionRepository().GetEvaluateArticle(Pet.Id)
		RouteArticleMap = make(map[string]*model.Article)
		RouteArticleMap["1"] = RouteArticle1
		RouteArticleMap["2"] = RouteArticle2
		PetEvaluateRouteArticleMap[Pet.Name] = RouteArticleMap
		return RouteArticleMap[Route]

	}

}

func (inst *petService) GetAllPets() []*model.Pet {
	Farm, _ := inst.GetFarmPets()
	Body, _ := inst.GetCarriedPetList()

	return append(Body, Farm...)
}

func (inst *petService) GetPet(PetName string) *model.Pet {
	Pets := inst.GetAllPets()
	for _, Pet := range Pets {
		if strings.Contains(Pet.Name, PetName) {
			return Pet
		}
	}
	return nil
}

//func (*petService) GetFarmPetByPetList(PetNameList []string) []*model.Pet {
//	pets, _ := repository.GetPetRepository().GetFarmedPets()
//	PetList := []*model.Pet{}
//	f
//	return PetList
//}
