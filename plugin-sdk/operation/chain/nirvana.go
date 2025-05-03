package chain

import (
	"context"
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/article"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/chain/spi/chain_model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/chain/spi/config"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/equip"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/pet"
	"strings"
	"time"
)

var NirvanaStateMachineInstance = &NirvanaStateMachineApi{}

type NirvanaStateMachineApi struct{}

var nirvanaCtx context.Context
var nirvanaCancel context.CancelFunc

func (NirvanaStateMachineApi) HandleFinishEvent() {
	nirvanaCancel()
}

func (NirvanaStateMachineApi) HandleStartEvent() {
	nirvanaCtx, nirvanaCancel = context.WithCancel(context.Background())
	go startNirvanaTask(nirvanaCtx)
}

func startNirvanaTask(context context.Context) {
	currentUser := status.GetLoginUser()
	GetNirvanaConfig := config.GetNirvanaConfig(currentUser.LoginName)
	NirvanaCount := 0
	NirvanaConfig := chain_model.CopyNirvanaConfig(*GetNirvanaConfig)
	for {
		if NirvanaCount >= GetNirvanaConfig.Finish.NirvanaCount && GetNirvanaConfig.Finish.NirvanaCount > 0 {
			plugin_log.Info("已涅槃 %d 次，涅槃数量达到预期 %d", NirvanaCount, GetNirvanaConfig.Finish.NirvanaCount)
			break
		}
		_, err := Nirvana(NirvanaConfig)
		if err != nil {
			return
		}
		NirvanaCount = NirvanaCount + 1
		Pet := pet.PetServiceInstance.GetBattlePet()
		plugin_log.Info("第 %d 次。涅槃成功 当前宠物 %s 成长: %f", NirvanaCount, Pet.Name, Pet.Cc)
	}
}

func Nirvana(Config *model.NirvanaConfig) (bool, error) {
	//
	startTime := time.Now()
	err1 := pet.PetServiceInstance.SaveUnBattlePet()
	if err1 != nil {
		return false, err1
	}
	// 获取主宠
	MainPet, err := getNirvanaPet(Config.MainPet)
	if err != nil {
		return false, err
	}
	err = pet.PetServiceInstance.SetBattlePet(MainPet.Id)
	if err != nil {
		return false, err
	}
	// 把其他宠物放到牧场
	err1 = pet.PetServiceInstance.SaveUnBattlePet()
	if err1 != nil {
		return false, err1
	}
	// 脱装备
	equip.EquipServiceImplInstance.OffEquip(MainPet.Id)
	// 获取副宠
	AtePet, err1 := getNirvanaPet(Config.AtePet)
	if err1 != nil {
		plugin_log.Error("获取副宠失败")
		return false, err1
	}
	// 获取捏
	NirvanaPet, err2 := getNirvanaPet(Config.NirvanaPet)
	if err2 != nil {
		plugin_log.Error("获取捏失败")
		return false, err1
	}
	once, err := nirvanaOnce(MainPet, AtePet, NirvanaPet, Config.ProtectType1, Config.ProtectType2)
	if err != nil {
		return false, err
	}
	minute, second := utils.CalculateTime(startTime)
	plugin_log.Info("涅槃成功")
	plugin_log.Info("本次耗时 %d 分钟 %d 秒", minute, second)
	return once, nil
}

func getNirvanaPet(config model.NirvanaPetConfig) (*model.Pet, error) {
	PetName := config.PetName
	if PetName == "" {
		return nil, errors.New("请输入宠物名称")
	}
	Pet := pet.PetServiceInstance.GetPet(PetName)
	if Pet == nil {
		if config.UseEgg {
			// 找到对应的蛋
			if config.IsNirvana {
				NirvanaEgg := operation.GetNirvanaEggArticle()
				if NirvanaEgg == nil {
					return nil, errors.New("找不到涅蛋")
				}
				err := article.ArticleServiceInstance.UserArticle(NirvanaEgg)
				if err != nil {
					return nil, err
				}
				// 偶发获取捏失败
				for Pet == nil {
					Pets, _ := pet.PetServiceInstance.GetCarriedPetList()
					for _, CarriedPet := range Pets {
						if !CarriedPet.IsBattle && strings.Contains(CarriedPet.Name, "兽") {
							Pet = CarriedPet
						}
					}
					if Pet == nil {
						time.Sleep(100 * time.Millisecond)
					}
				}

			} else {
				if PetName == "小神龙" {
					GOD_EGG := operation.GetDragonEggArticle()
					if GOD_EGG == nil {
						return nil, errors.New("找不到小神蛋")
					}
					err := article.ArticleServiceInstance.UserArticle(GOD_EGG)
					if err != nil {
						return nil, err
					}
					// 偶发获取捏失败
					for Pet == nil {
						Pets, _ := pet.PetServiceInstance.GetCarriedPetList()
						for _, CarriedPet := range Pets {
							if !CarriedPet.IsBattle && strings.Contains(CarriedPet.Name, "小神龙") {
								Pet = CarriedPet
							}
						}
						if Pet == nil {
							time.Sleep(100 * time.Millisecond)
						}
					}
				} else {
					ArticleList, err := article.ArticleServiceInstance.QueryArticleList(PetName)
					if err != nil {
						return nil, err
					}
					for _, Article := range ArticleList {
						if Article.ArticleType == "宠物卵" {
							err = article.ArticleServiceInstance.UserArticle(Article)
							if err != nil {
								return nil, err
							}
							break
						}
					}
					Pet = pet.PetServiceInstance.GetPet(PetName)
				}

			}
		} else {
			return nil, errors.New("未找到宠物")
		}
	}
	if Pet == nil {
		return nil, errors.New("未找到宠物")
	}
	// 进化
	err := prepareNirvana(Pet, config)
	if err != nil {
		return nil, err
	}
	return Pet, nil
}

func prepareNirvana(Pet *model.Pet, config model.NirvanaPetConfig) error {
	TempPet, _ := pet.PetServiceInstance.GetPetDetail(Pet.Id)
	Pet.Level = TempPet.Level
	// 升级
	err := pet.PetServiceInstance.SetBattlePet(Pet.Id)
	if err != nil {
		return err
	}
	if config.PetLevel < 60 {
		config.PetLevel = 60
	}
	ExperienceType := config.ExperienceType
	ExperienceArticles, err := operation.GetExperienceArticleByType(ExperienceType)
	//ArticleServiceInstance.QueryArticleListByNameLists(ExperienceNameList)
	if err != nil {
		plugin_log.Error("获取经验失败, %s", err.Error())
		return err
	}
	if ExperienceArticles == nil {
		return errors.New("物品不足： " + ExperienceType)
	}
	for {
		if Pet.Level >= config.PetLevel {
			break
		}
		plugin_log.Info("开始吃经验")
		expArticle := ExperienceArticles
		if expArticle.ArticleCount == 0 {
			continue
		}
		article.ArticleServiceInstance.UserArticle(expArticle)
		CurrentPetStatus, getPetError := pet.PetServiceInstance.GetPetDetail(Pet.Id)
		if getPetError != nil {
			CurrentPetStatus, _ = pet.PetServiceInstance.GetPetDetail(Pet.Id)
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
			plugin_log.Error("进化失败 退出进化")
			break
		}
		CurrentPetStatus, getPetError := pet.PetServiceInstance.GetPetDetail(Pet.Id)
		if getPetError != nil {
			CurrentPetStatus, _ = pet.PetServiceInstance.GetPetDetail(Pet.Id)
		}
		if CurrentPetStatus.Cc >= config.PetCc && config.PetCc > 0 {
			Pet.Cc = CurrentPetStatus.Cc
			Pet.Name = CurrentPetStatus.Name
			plugin_log.Info("当前宠物: %s 成长为 %f,达到cc阈值: %f", Pet.Name, Pet.Cc, config.PetCc)
			return errors.New("CC达到阈值")
		}
	}
	return nil
}

func nirvanaOnce(MainPet *model.Pet, AtePet *model.Pet, NirvanaPet *model.Pet,
	ArticleType1, ArticleType2 string) (bool, error) {
	ProtectArticle1, err1 := operation.GetProtectArticleByType(ArticleType1)
	if err1 != nil {
		return false, err1
	}
	ProtectArticle2, err2 := operation.GetProtectArticleByType(ArticleType2)
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
