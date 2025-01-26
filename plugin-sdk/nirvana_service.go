package plugin_sdk

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"time"
)

var NirvanaServiceImplInstance = &nirvanaServiceImpl{}

type nirvanaServiceImpl struct {
}

func (inst *nirvanaServiceImpl) Nirvana(Config *model.NirvanaConfig) (bool, error) {
	//
	startTime := time.Now()
	err1 := PetServiceInstance.SaveUnBattlePet()
	if err1 != nil {
		return false, err1
	}
	// 获取主宠
	MainPet, _ := getNirvanaPet(Config.MainPet)
	err := PetServiceInstance.SetBattlePet(MainPet.Id)
	if err != nil {
		return false, err
	}
	// 把其他宠物放到牧场
	err1 = PetServiceInstance.SaveUnBattlePet()
	if err1 != nil {
		return false, err1
	}
	// 脱装备
	EquipServiceImplInstance.OffEquip(MainPet.Id)
	// 获取副宠
	AtePet, _ := getNirvanaPet(Config.AtePet)
	// 获取捏
	NirvanaPet, _ := getNirvanaPet(Config.NirvanaPet)
	once, err := nirvanaOnce(MainPet, AtePet, NirvanaPet, Config.ProtectType1, Config.ProtectType2)
	if err != nil {
		return false, err
	}
	minute, second := utils.CalculateTime(startTime)
	log.Info("涅槃成功")
	log.Info("本次耗时 %d 分钟 %d 秒", minute, second)
	return once, nil
}

func getNirvanaPet(config model.NirvanaPetConfig) (*model.Pet, error) {
	PetName := config.PetName
	if PetName == "" {
		return nil, errors.New("请输入宠物名称")
	}
	Pet := PetServiceInstance.GetPet(PetName)
	if Pet == nil {
		if config.UseEgg {
			// 找到对应的蛋
			if config.IsNirvana {
				NirvanaEgg := GetNirvanaEggArticle()
				if NirvanaEgg == nil {
					return nil, errors.New("找不到涅蛋")
				}
			} else {
				ArticleList, err := ArticleServiceInstance.QueryArticleList(PetName)
				if err != nil {
					return nil, err
				}
				for _, Article := range ArticleList {
					if Article.ArticleType == "宠物卵" {
						err = ArticleServiceInstance.UserArticle(Article)
						if err != nil {
							return nil, err
						}
						break
					}
				}
				Pet = PetServiceInstance.GetPet(PetName)
			}
		} else {
			return nil, errors.New("")
		}
	}

	// 进化
	err := prepareNirvana(Pet, config)
	if err != nil {
		return nil, err
	}
	return Pet, nil
}

func prepareNirvana(Pet *model.Pet, config model.NirvanaPetConfig) error {
	// 升级
	err := PetServiceInstance.SetBattlePet(Pet.Id)
	if err != nil {
		return err
	}
	if config.PetLevel < 60 {
		config.PetLevel = 60
	}
	ExperienceType := config.ExperienceType
	ExperienceArticles, err := GetExperienceArticleByType(ExperienceType)
	//ArticleServiceInstance.QueryArticleListByNameLists(ExperienceNameList)
	if err != nil {
		log.Error("获取经验失败, %s", err.Error())
		return err
	}
	if ExperienceArticles == nil {
		return errors.New("物品不足： " + ExperienceType)
	}
	for {
		if Pet.Level >= config.PetLevel {
			break
		}
		log.Info("开始吃经验")
		article := ExperienceArticles
		if article.ArticleCount == 0 {
			continue
		}
		ArticleServiceInstance.UserArticle(article)
		CurrentPetStatus, getPetError := PetServiceInstance.GetPetDetail(Pet.Id)
		if getPetError != nil {
			CurrentPetStatus, _ = PetServiceInstance.GetPetDetail(Pet.Id)
		}
		Pet.Level = CurrentPetStatus.Level
		Pet.Name = CurrentPetStatus.Name
		if CurrentPetStatus.Level >= config.PetLevel {
			break
		}
	}

	// 进化
	for _, EvaluateConfigInstance := range config.Evaluate {
		evaluate, evaluateErr := Evaluate(Pet, EvaluateConfigInstance.EvaluateRoute)
		if evaluateErr != nil || evaluate == false {
			if EvaluateConfigInstance.ForceEvaluate {
				return err
			}
			log.Error("进化失败 退出进化")
			break
		}
		CurrentPetStatus, getPetError := PetServiceInstance.GetPetDetail(Pet.Id)
		if getPetError != nil {
			CurrentPetStatus, _ = PetServiceInstance.GetPetDetail(Pet.Id)
		}
		if CurrentPetStatus.Cc >= config.PetCc && config.PetCc > 0 {
			Pet.Cc = CurrentPetStatus.Cc
			Pet.Name = CurrentPetStatus.Name
			log.Info("当前宠物: %s 成长为 %f,达到cc阈值: %f", Pet.Name, Pet.Cc, config.PetCc)
			return errors.New("CC达到阈值")
		}
	}
	return nil
}

func nirvanaOnce(MainPet *model.Pet, AtePet *model.Pet, NirvanaPet *model.Pet,
	ArticleType1, ArticleType2 string) (bool, error) {
	ProtectArticle1, err1 := GetProtectArticleByType(ArticleType1)
	if err1 != nil {
		return false, err1
	}
	ProtectArticle2, err2 := GetProtectArticleByType(ArticleType2)
	if err2 != nil {
		return false, err2
	}
	result, err := repository.GetFusionRepository().Nirvana(*MainPet, *AtePet, *NirvanaPet, *ProtectArticle1, *ProtectArticle2)
	if err != nil {
		return false, err
	}
	if !result {
		return nirvanaOnce(MainPet, AtePet, NirvanaPet, ArticleType1, ArticleType2)
	}

	return true, nil
}
