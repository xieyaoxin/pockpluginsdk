package plugin_sdk

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"time"
)

var FusionServiceImplInstance = &fusionServiceImpl{}

type fusionServiceImpl struct {
}

// Fusion 合成一次
func Fusion(mergeConfig model.SingleMergeConfig, MainForceEvaluate bool, FusionAfterCcThreshold bool, AteForceEvaluate bool) (bool, error) {
	Pet1, petError1 := getFusionPet(&mergeConfig.MainPetConfig)
	if petError1 != nil {
		return false, petError1
	}
	petError1 = prepareMerge(Pet1, &mergeConfig.MainPetConfig, MainForceEvaluate)
	if petError1 != nil {
		if petError1.Error() == "CC达到阈值" {
			// 主宠CC到达阈值后是否再合一次， 不合直接退出
			if !FusionAfterCcThreshold {
				return true, nil
			}
		} else {
			return false, petError1
		}
	}
	Pet2, petError2 := getFusionPet(&mergeConfig.AtePetConfig)
	if petError2 != nil {
		return false, petError2
	}
	petError2 = prepareMerge(Pet2, &mergeConfig.AtePetConfig, AteForceEvaluate)
	if petError2 != nil && petError2.Error() != "CC达到阈值" {
		return false, petError2
	}
	// todo 增加宠物检查
	ProtectArticle1, err1 := GetProtectArticleByType(mergeConfig.ProtectType1)
	if err1 != nil {
		return false, err1
	}
	ProtectArticle2, err2 := GetProtectArticleByType(mergeConfig.ProtectType2)
	if err2 != nil {
		return false, err2
	}
	response, err3 := repository.GetFusionRepository().Fusion(*Pet1, *Pet2, *ProtectArticle1, *ProtectArticle2)
	if err3 != nil {
		return false, err3
	}
	if response {
		return true, nil
	} else {
		log.Info("合成失败，等待2秒")
		time.Sleep(2 * time.Second)
		return Fusion(mergeConfig, MainForceEvaluate, FusionAfterCcThreshold, AteForceEvaluate)
	}
}

func Evaluate(Pet *model.Pet, EvaluateRoute string) (bool, error) {
	EvaluateArticle := PetServiceInstance.GetPetEvaluateArticle(Pet, EvaluateRoute)
	if EvaluateArticle == nil {
		return false, errors.New("获取进化物品失败")
	}
	return repository.GetFusionRepository().Evaluate(Pet, EvaluateRoute, EvaluateArticle.Pid, "0")

}
func MergeGod(Config model.MergeGodConfig) (bool, error) {
	startTime := time.Now()
	// 寄存所有非主站宠物
	PetServiceInstance.SaveUnBattlePet()
	//
	MainPet, err1 := mergeDragon(Config.MainPet)
	if err1 != nil {
		log.Info("合成失败 %s", err1.Error())
		return false, err1
	}
	AteDragon, err3 := mergeDragon(Config.AteDragon)
	if err3 != nil {
		return false, err3
	}
	EatDragonConfig := model.CopySingleMergeConfig(*Config.EatDragon)
	EatDragonConfig.MainPetConfig.PetId = MainPet.Id
	EatDragonConfig.AtePetConfig.PetId = AteDragon.Id
	_, err := Fusion(EatDragonConfig, false, true, false)
	if err != nil {
		return false, err
	}
	FusionResult := PetServiceInstance.GetBattlePet()
	log.Info("单次合神结束,合成结果: %s 成长 %f ", FusionResult.Name, FusionResult.Cc)
	minute, second := utils.CalculateTime(startTime)
	log.Info("本次耗时 %d 分钟 %d 秒", minute, second)
	return FusionResult.Name == "小神龙琅玡" || FusionResult.Name == "白虎", nil

}

func mergeDragon(Config *model.MergeDragonConfig) (*model.Pet, error) {
	MainPet, err1 := mergeWithPetType(Config.MainPet)
	if err1 != nil {
		return nil, err1
	}
	if Config.AtePet == nil {
		return MainPet, nil
	}
	// 合副宠 吃副宠
	EatConfig := model.CopySingleMergeConfig(Config.EatPet)
	for MainPet.Cc < Config.EatPet.MainPetConfig.PetCc {
		// 针对部分宠物进行特殊处理
		for utils.SlicesContainsString(FusionSpecialEvaluatePetNames, MainPet.Name) {
			// 吃一个册子
			//ExperienceArticle := GetExperienceArticleByType(EatConfig.MainPetConfig.ExperienceType)
			TempMergeConfig := model.CopyMergePetConfig(EatConfig.MainPetConfig)
			TempMergeConfig.Evaluate = []*model.EvaluateConfig{
				&model.EvaluateConfig{
					EvaluateRoute: "2",
				},
			}
			err := prepareMerge(MainPet, &TempMergeConfig, true)
			if err != nil {
				return nil, err
			}
			TempPet, _ := PetServiceInstance.GetPetDetail(MainPet.Id)
			MainPet.Name = TempPet.Name
			MainPet.Cc = TempPet.Cc

		}
		AtePet, err2 := mergeWithPetType(Config.AtePet)
		if err2 != nil {
			return nil, err2
		}
		EatConfig.MainPetConfig.PetId = MainPet.Id
		EatConfig.AtePetConfig.PetId = AtePet.Id
		fusion, err := Fusion(EatConfig, false, false, false)
		if fusion && err == nil {
			CurrentMainPet := PetServiceInstance.GetBattlePet()
			MainPet.Id = CurrentMainPet.Id
			MainPet.Name = CurrentMainPet.Name
			MainPet.Cc = CurrentMainPet.Cc
		} else {
			return nil, err
		}

	}
	return MainPet, nil
}

// mergeWithPetType 根据宠物类型合成主宠/副宠/副龙
func mergeWithPetType(config *model.SingleMergeConfig) (*model.Pet, error) {
	// 存其他宠物
	PetServiceInstance.SaveUnBattlePet()
	PetType := config.MainPetConfig.PetType
	switch PetType {
	case "LD":
		WXL := GetPetFromCacheByPetName(WXL_NAMES)
		if WXL != nil {
			return WXL, nil
		}
		// 先读缓存
		if len(LD_EGG_ARTICLE_LIST) == 0 {
			return nil, errors.New("找不到龙蛋")
		}
		// 使用龙蛋
		for _, LD := range LD_EGG_ARTICLE_LIST {
			if LD.ArticleCount > 0 {
				ArticleServiceInstance.UserArticle(LD)
				break
			}
		}
		// 检查身上的宠物
		list, _ := PetServiceInstance.GetCarriedPetList()
		if len(list) > 1 {
			for _, pet := range list {
				if !pet.IsBattle {
					return pet, nil
				}
			}
		} else {
			return nil, errors.New("未找到龙")
		}
		break
	case "WX":
		// todo 根据中间产物获取不同的进化路线
		WxMergeConfig := initWXDefaultMergeConfig()
		BmMergeConfig := initAteBmDefaultMergeConfig()
		fusion, err := Fusion(model.SingleMergeConfig{
			MainPetConfig: WxMergeConfig,
			AtePetConfig:  BmMergeConfig,
			ProtectType1:  config.ProtectType1,
			ProtectType2:  config.ProtectType2,
		}, true, true, false)
		if err != nil {
			return nil, err
		}
		if fusion {
			BattlePet := PetServiceInstance.GetBattlePet()
			return BattlePet, err
		} else {
			return nil, errors.New("合成失败")
		}
	case "ALY":
		break
	case "BMW":
		BmMergeConfig1 := initAteBmDefaultMergeConfig()
		BmMergeConfig2 := initAteBmDefaultMergeConfig()
		fusion, err := Fusion(model.SingleMergeConfig{
			MainPetConfig: BmMergeConfig1,
			AtePetConfig:  BmMergeConfig2,
			ProtectType1:  config.ProtectType1,
			ProtectType2:  config.ProtectType2,
		}, true, true, false)
		if err != nil {
			return nil, err
		}
		if fusion {
			BattlePet := PetServiceInstance.GetBattlePet()
			return BattlePet, err
		} else {
			return nil, errors.New("合成失败")
		}
	case "BM":
		break
	}
	return nil, errors.New("未知错误")
}

func prepareMerge(Pet *model.Pet, config *model.MergePetConfig, ForceEvaluate bool) error {
	// 设置主站
	PetServiceInstance.SetBattlePet(Pet.Id)
	// 吃经验
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
		if CurrentPetStatus.Level >= config.PetLevel {
			Pet.Level = CurrentPetStatus.Level
			Pet.Name = CurrentPetStatus.Name
			break
		}
	}
	// 开始进化
	if config.Evaluate == nil {
		return nil
	}
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
		if CurrentPetStatus.Cc >= config.PetCc && config.PetCc > 0 && !ForceEvaluate {
			Pet.Cc = CurrentPetStatus.Cc
			Pet.Name = CurrentPetStatus.Name
			log.Info("当前宠物: %s 成长为 %f,达到cc阈值: %f", Pet.Name, Pet.Cc, config.PetCc)
			return errors.New("CC达到阈值")
		}
	}
	return nil
}

func getFusionPet(Config *model.MergePetConfig) (*model.Pet, error) {
	if Config.PetId == "" {
		// 根据PetType抓宠
		BM := GetBMFromCache()
		Config.PetId = BM.Id
		// 根据PetType定义进化路线
		if Config.Evaluate == nil && Config.PetType == "WX" {
			EvaluateConfigs := initWXDefaultEvaluate(BM.Name)
			Config.Evaluate = EvaluateConfigs
		}
	}
	Pet, _ := PetServiceInstance.GetPetDetail(Config.PetId)
	//err := prepareMerge(Pet, &Config)
	//if err != nil {
	//	return nil, err
	//}
	//Pet, _ = PetServiceInstance.GetPetDetail(Config.PetId)
	return Pet, nil
}

// 获取五系默认的配置
func initWXDefaultMergeConfig() model.MergePetConfig {
	return model.MergePetConfig{
		PetType:        "WX",
		ExperienceType: "1E",
		PetLevel:       60,
	}
}

// // 获取五系配置
//func getWx() (*model.Pet, model.MergePetConfig, error) {
//	BM := GetBMFromCache()
//	// 根据不同的波姆获取进化路线
//	EvaluateConfigs := initWXDefaultEvaluate(BM.Name)
//	MergeConfig := initWXDefaultMergeConfig()
//	MergeConfig.PetId = BM.Id
//	MergeConfig.Evaluate = EvaluateConfigs
//	err := prepareMerge(BM, &MergeConfig)
//	if err != nil {
//		return nil, MergeConfig, err
//	}
//	return nil, MergeConfig, err
//}

//// 捞一只狗粮 进化一次
//func getAteBm() (*model.Pet, model.MergePetConfig, error) {
//	BM := GetBMFromCache()
//	MergeConfig := initAteBmDefaultMergeConfig()
//	MergeConfig.PetId = BM.Id
//	err := prepareMerge(BM, &MergeConfig)
//	if err != nil {
//		return nil, MergeConfig, err
//	}
//	return BM, MergeConfig, nil
//}

// 获取副宠波姆默认的配置
func initAteBmDefaultMergeConfig() model.MergePetConfig {
	return model.MergePetConfig{
		PetType:        "BM",
		ExperienceType: "1E",
		PetLevel:       60,
		Evaluate: []*model.EvaluateConfig{
			&model.EvaluateConfig{
				EvaluateRoute: "1",
				ForceEvaluate: false,
			},
		},
	}
}

func initWXDefaultEvaluate(BmName string) []*model.EvaluateConfig {
	switch BmName {
	case "火波姆":
		return []*model.EvaluateConfig{
			// 高级进化之书	 火光球	进化之书	五色羽毛	高级进化之书	 蝙蝠翅膀	火灵猴的桃子
			&model.EvaluateConfig{"2", true}, &model.EvaluateConfig{"1", true},
			&model.EvaluateConfig{"1", true}, &model.EvaluateConfig{"1", true},
			&model.EvaluateConfig{"1", true}, &model.EvaluateConfig{"2", false},
			&model.EvaluateConfig{"2", true},
		}
	case "金波姆":
		return []*model.EvaluateConfig{
			// 高级进化之书	雷光球 超级进化书 紫貘之珠	超级进化书
			&model.EvaluateConfig{"2", true}, &model.EvaluateConfig{"1", true},
			&model.EvaluateConfig{"2", true}, &model.EvaluateConfig{"2", true},
			&model.EvaluateConfig{"1", true},
		}
	case "绿波姆":
		return []*model.EvaluateConfig{
			// 进化之书  超级进化之书	进化之书	进化之书	青龙珠 -青蛟之魄/超级进化之书
			&model.EvaluateConfig{"1", true}, &model.EvaluateConfig{"2", true},
			&model.EvaluateConfig{"1", true}, &model.EvaluateConfig{"1", true},
			&model.EvaluateConfig{"2", true}, &model.EvaluateConfig{"2", false},
			&model.EvaluateConfig{"1", false},
		}
	case "水波姆":
		// 进化之书 超级进化书 超级进化书 妖精之泪
		return []*model.EvaluateConfig{
			&model.EvaluateConfig{"1", true}, &model.EvaluateConfig{"2", true},
			&model.EvaluateConfig{"2", true}, &model.EvaluateConfig{"2", true},
		}
	case "土波姆":
		// 高级进化之书四叶草	月亮水晶	战神之盔
		return []*model.EvaluateConfig{
			&model.EvaluateConfig{"2", true}, &model.EvaluateConfig{"2", true},
			&model.EvaluateConfig{"2", true}, &model.EvaluateConfig{"2", true},
		}
	}
	return nil
}
