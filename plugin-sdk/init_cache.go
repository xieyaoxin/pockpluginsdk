package plugin_sdk

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
)

var LD_EGG_ARTICLE_LIST = []*model.Article{}
var EXPERIENCE_TYPE_ARTICEL_MAP = make(map[string][]*model.Article)
var PROTECT_ARTICEL_MAP = make(map[string][]*model.Article)
var FARM_PETS = []*model.Pet{}
var BM_NAMES = []string{"金波姆", "绿波姆", "水波姆", "火波姆", "土波姆"}
var BMW_NAMES = []string{"金波姆王", "绿波姆王", "水波姆王", "火波姆王", "土波姆王", "碧蟾", "魔岩卵", "水仙", "火芒", "金光鼠"}
var WXL_NAMES = []string{"青龙琅琅", "小青龙琅琅", "金龙霸王", "冰龙苍海", "艾薇儿", "三尾忍忍", "仙狐六尾", "天狐莫姬", "炎龙血焰", "黄龙莫虚"}
var DRAGON_EGG_NAME = []string{"青龙琅琅之卵", "小青龙琅琅", "金龙霸王之卵", "冰龙苍海之卵", "艾薇儿之卵", "三尾忍忍", "仙狐六尾", "天狐莫姬之卵", "炎龙血焰之卵", "黄龙莫虚之卵"}
var PetEvaluateRouteArticleMap = make(map[string]map[string]*model.Article)
var FusionSpecialEvaluatePetNames = []string{"三尾忍忍", "狐仙六尾", "小青龙琅琅"}
var NIRVANA_EGG_LIST = []*model.Article{}

// 小神蛋
var DRAGON_EGG_LIST = []*model.Article{}

func init() {
	//InitMergeArticleCache()
}

func InitMergeArticleCache() {
	// 缓存经验类物品信息
	for _, ExperienceType := range repository.GetFusionRepository().GetExperienceTypeList() {
		ArticleNameList := repository.GetFusionRepository().GetExperienceList(ExperienceType)
		ArticleList, _ := ArticleServiceInstance.QueryArticleListByNameLists(ArticleNameList)
		EXPERIENCE_TYPE_ARTICEL_MAP[ExperienceType] = ArticleList
	}
	// 缓存龙蛋信息
	LD_EGG_ARTICLE_LIST, _ = ArticleServiceInstance.QueryArticleListByNameLists(DRAGON_EGG_NAME)
	// 缓存保护石信息
	for _, ProtectArticleType := range repository.GetFusionRepository().GetProtectArticleTypeList() {
		ArticleNameList := repository.GetFusionRepository().GetProjectArticleList(ProtectArticleType)
		ArticleList, _ := ArticleServiceInstance.QueryArticleListByNameLists(ArticleNameList)
		PROTECT_ARTICEL_MAP[ProtectArticleType] = ArticleList
	}
	// 缓存宠物信息
	FARM_PETS = PetServiceInstance.GetAllPets()
}

func InitNirvanaCache() {
	for _, ExperienceType := range repository.GetFusionRepository().GetExperienceTypeList() {
		ArticleNameList := repository.GetFusionRepository().GetExperienceList(ExperienceType)
		ArticleList, _ := ArticleServiceInstance.QueryArticleListByNameLists(ArticleNameList)
		EXPERIENCE_TYPE_ARTICEL_MAP[ExperienceType] = ArticleList
	}

	for _, ProtectArticleType := range repository.GetFusionRepository().GetNirvanaArticleTypeList() {
		ArticleNameList := repository.GetFusionRepository().GetNirvanaArticleList(ProtectArticleType)
		ArticleList, _ := ArticleServiceInstance.QueryArticleListByNameLists(ArticleNameList)
		PROTECT_ARTICEL_MAP[ProtectArticleType] = ArticleList
	}

	NIRVANA_EGG_LIST, _ = ArticleServiceInstance.QueryArticleListByNameLists(model.NirvanaEggList)
	DRAGON_EGG_LIST, _ = ArticleServiceInstance.QueryArticleListByNameLists(model.DragonEggList)
	FARM_PETS = PetServiceInstance.GetAllPets()
}

// GetBMFromCache 获取波姆, 先从缓存中获取,缓存中没有则去捕捉
func GetBMFromCache() *model.Pet {
	BM := GetPetFromCacheByPetName(BM_NAMES)
	if BM != nil {
		return BM
	}
	err := PetServiceInstance.SaveUnBattlePet()
	if err != nil {
		plugin_log.Error("寄存非主站宠物失败")
		return nil
	}
	catchBmConfig := InitCatchBmConfig()
	for {
		result := FightOneTime(*catchBmConfig)
		if result == "捕捉成功" {
			break
		}
	}
	PetCarried, _ := PetServiceInstance.GetCarriedPetList()
	for _, Pet := range PetCarried {
		if !Pet.IsBattle {
			return Pet
		}
	}
	return nil
}

// GetBMFromCache 获取波姆, 先从缓存中获取,缓存中没有则去捕捉
func GetBMWFromCache() *model.Pet {
	BMW := GetPetFromCacheByPetName(BMW_NAMES)
	return BMW
}

func GetProtectArticleByType(ProtectType string) (*model.Article, error) {
	if ProtectArticleList, exists := PROTECT_ARTICEL_MAP[ProtectType]; exists {
		TempProtectArticle := make([]*model.Article, len(ProtectArticleList))
		copy(TempProtectArticle, ProtectArticleList)
		for _, ProtectArticle := range TempProtectArticle {
			if ProtectArticle.ArticleCount == 0 {
				PROTECT_ARTICEL_MAP[ProtectType] = model.ArticleSliceRemoveItem(ProtectArticleList, ProtectArticle)
				continue
			} else {
				ProtectArticle.ArticleCount = ProtectArticle.ArticleCount - 1
				return ProtectArticle, nil
			}
		}
	}
	// todo 增加从背包获取物品
	return nil, errors.New("找不到合宠物品: " + ProtectType)
}
func GetPetFromCacheByPetName(PetNameList []string) *model.Pet {
	result := []*model.Pet{}
	var tempPet *model.Pet
	for _, Pet := range FARM_PETS {
		if util.SlicesContainsString(PetNameList, Pet.Name) {
			tempPet = Pet
			break
		} else {
			result = append(result, Pet)
		}
	}
	FARM_PETS = model.PetSliceRemoveItem(FARM_PETS, tempPet)
	return tempPet
}
func GetExperienceArticleByType(ExperienceType string) (*model.Article, error) {
	if ExperienceTypeList, exists := EXPERIENCE_TYPE_ARTICEL_MAP[ExperienceType]; exists {
		TempExperienceArticle := make([]*model.Article, len(ExperienceTypeList))
		copy(TempExperienceArticle, ExperienceTypeList)
		for _, ExperienceArticle := range TempExperienceArticle {
			if ExperienceArticle.ArticleCount == 0 {
				PROTECT_ARTICEL_MAP[ExperienceType] = model.ArticleSliceRemoveItem(ExperienceTypeList, ExperienceArticle)
				continue
			} else {
				ExperienceArticle.ArticleCount = ExperienceArticle.ArticleCount - 1
				return ExperienceArticle, nil
			}
		}
	}
	// todo 增加从背包获取物品
	return nil, errors.New("找不到合宠物品: " + ExperienceType)
}
func GetNirvanaEggArticle() *model.Article {
	TempNirvanaArticle := make([]*model.Article, len(NIRVANA_EGG_LIST))
	copy(TempNirvanaArticle, NIRVANA_EGG_LIST)
	for _, NirvanaArticle := range TempNirvanaArticle {
		if NirvanaArticle.ArticleCount == 0 {
			NIRVANA_EGG_LIST = model.ArticleSliceRemoveItem(NIRVANA_EGG_LIST, NirvanaArticle)
			continue
		} else {
			NirvanaArticle.ArticleCount = NirvanaArticle.ArticleCount - 1
			return NirvanaArticle
		}
	}
	return nil
}

func GetDragonEggArticle() *model.Article {
	TempDUArticle := make([]*model.Article, len(DRAGON_EGG_LIST))
	copy(TempDUArticle, DRAGON_EGG_LIST)
	for _, DragonArticle := range TempDUArticle {
		if DragonArticle.ArticleCount == 0 {
			DRAGON_EGG_LIST = model.ArticleSliceRemoveItem(DRAGON_EGG_LIST, DragonArticle)
			continue
		} else {
			DragonArticle.ArticleCount = DragonArticle.ArticleCount - 1
			return DragonArticle
		}
	}
	return nil
}
