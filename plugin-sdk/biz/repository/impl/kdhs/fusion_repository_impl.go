package kdhs

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"strings"
	"time"
)

var FusionRepositorykdhsImplInstance = &fusionRepositoryKdhsImpl{}

// 经验类型

var experienceTypeDict = make(map[string][]string)

var protectArticleDict = make(map[string][]string)

var nirvanaArticleDict = make(map[string][]string)

func init() {
	// 初始化经验字典
	experienceTypeDict[model.EXPERIENCE_200W] = []string{"200w经验月饼"}
	experienceTypeDict[model.EXPERIENCE_500W] = []string{}
	experienceTypeDict[model.EXPERIENCE_5000W] = []string{}
	experienceTypeDict[model.EXPERIENCE_1E] = []string{"修炼仙册(限时)", "修炼仙册"}
	experienceTypeDict[model.EXPERIENCE_10E] = []string{"10亿经验卷轴"}
	experienceTypeDict[model.EXPERIENCE_20E] = []string{"20亿经验卷轴"}
	experienceTypeDict[model.EXPERIENCE_200E] = []string{"200亿经验卷轴(限时)", "15亿经验卷轴"}
	experienceTypeDict[model.EXPERIENCE_800E] = []string{"15亿经验卷轴", "800亿经验卷轴"}
	// 初始化保护石字典
	protectArticleDict[model.PROTECT_HC] = []string{}
	protectArticleDict[model.PROTECT_SH] = []string{}
	protectArticleDict[model.PROTECT_TS] = []string{"天神之石"}
	protectArticleDict[model.PROTECT_ZZ] = []string{"至尊神石(绑定)", "至尊神石"}
	protectArticleDict[model.PROTECT_3XCC] = []string{"★★★成长魂石【绑定】", "★★★成长魂石"}
	protectArticleDict[model.PROTECT_3X] = []string{}

	nirvanaArticleDict[model.PROTECT_NIRVANA_SD] = []string{"涅盘神丹(限时)", "涅盘神丹(绑定)", "涅盘神丹"}
	nirvanaArticleDict[model.PROTECT_NIRVANA_NPD] = []string{"涅盘丹"}
	nirvanaArticleDict[model.PROTECT_NIRVANA_SSD] = []string{"涅盘圣丹"}

	nirvanaArticleDict[model.PROTECT_NIRVANA_YPNC] = []string{"一品捏成丹"}
	nirvanaArticleDict[model.PROTECT_NIRVANA_SPNC] = []string{"上品捏成丹"}
	nirvanaArticleDict[model.PROTECT_NIRVANA_JPNC] = []string{"极品捏成丹"}
	nirvanaArticleDict[model.PROTECT_NIRVANA_SPMZ] = []string{"上品命中丹"}
	nirvanaArticleDict[model.PROTECT_NIRVANA_JPMZ] = []string{"极品命中丹"}
	nirvanaArticleDict[model.PROTECT_NIRVANA_JPNG] = []string{"极品涅攻丹"}
}

type fusionRepositoryKdhsImpl struct{}

func (inst *fusionRepositoryKdhsImpl) Fusion(Pet1 model.Pet, Pet2 model.Pet, protect1 model.Article, protect2 model.Article) (bool, error) {
	params := util.InitParam()
	params["ap"] = Pet1.Id
	params["bp"] = Pet2.Id
	params["p1"] = protect1.ID
	params["p2"] = protect2.ID
	params["type1"] = "check"
	mergeResult := CallServerGetInterface("function/cmpGate.php", params)
	log.Info("合成结果 主宠%s, 副宠%s 结果 %s", Pet1.Name, Pet2.Name, mergeResult)
	switch mergeResult {
	case "0", "6":
		log.Error("合成失败")
		return false, nil
	case "1":
		log.Error("找不到对应宠物")
		return false, errors.New("找不到对应宠物")
	case "2":
		log.Error("这两个宠物好像不能合成噢")
		return false, errors.New("这两个宠物好像不能合成噢")
	case "3":
		log.Error("您的金币不足，不能合成")
		return false, errors.New("您的金币不足，不能合成")
	case "20":
		log.Error("对不起，您的背包中没有此物品!")
		return false, errors.New("找不到对应宠物")
	case "5":
		log.Info("恭喜你，合成成功")
		return true, nil
	case "10":
		log.Error("数据读取失败")
		return false, errors.New("数据读取失败")
	case "11":
		log.Error("冷却时间未到")
		time.Sleep(1 * time.Second)
		return inst.Fusion(Pet1, Pet2, protect1, protect2)
	case "15":
		log.Error("宠物成长未达到合成条件哦")
		return false, errors.New("宠物成长未达到合成条件哦")
	default:
		log.Error("找不到对应宠物")
		return false, nil
	}
}

func (inst *fusionRepositoryKdhsImpl) Nirvana(Pet1 model.Pet, Pet2 model.Pet, NPS model.Pet, protect1 model.Article, protect2 model.Article) (bool, error) {
	params := util.InitParam()
	params["ap"] = Pet1.Id
	params["bp"] = Pet2.Id
	params["zs"] = NPS.Id
	params["p1"] = protect1.ID
	params["p2"] = protect2.ID
	result := CallServerGetInterface("function/zsGate.php", params)
	//
	//if result == "11" {
	//	property.Log.Printf("涅磐结果 主宠%s, 副宠%s 结果 %s", util.MapToJsonString(mainPet), util.MapToJsonString(atePet), result)
	//	return Nirvana(property, mainPet, atePet, nirvana, article1, article2)
	//}
	//property.Log.Printf("涅磐结果 主宠%s, 副宠%s 结果 %s", util.MapToJsonString(mainPet), util.MapToJsonString(atePet), result)
	switch result {
	case "0":
		log.Error("转生失败！")
		return false, nil
	case "1":
		log.Error("您没有相应的宠物！")
		return false, errors.New("您没有相应的宠物！")
	case "2":
		log.Error("这两个宠物好像不能转生噢！")
		return false, errors.New("这两个宠物好像不能转生噢！")
	case "3":
		log.Error("您的金币不足，不能转生！")
		return false, errors.New("您的金币不足，不能转生！")
	case "7":
		log.Error("请选择涅磐兽！")
		return false, errors.New("请选择涅磐兽！")
	case "6":
		log.Error("对不起，转生失败！")
		return false, nil
	case "5":
		log.Info("恭喜你，转生成功!！")
		return true, nil
	case "10":
		log.Error("数据读取失败!！")
		return false, errors.New("数据读取失败!！")
	case "11":
		log.Error("冷却时间未到！")
		return false, errors.New("冷却时间未到！！")
	default:
		log.Error("转生失败！")
		return inst.Nirvana(Pet1, Pet2, NPS, protect1, protect2)
	}
}

func (inst *fusionRepositoryKdhsImpl) GetEvaluateArticle(PetId string) (*model.Article, *model.Article) {
	params := util.InitParam()
	params["pid"] = PetId
	result := CallServerGetInterface("function/Sd_Mod.php", params)
	//	 8178204
	lines := strings.Split(result, "\n")
	EvaulateArticleRoute1 := &model.Article{}
	EvaulateArticleRoute2 := &model.Article{}
	for index, line := range lines {
		if strings.Contains(line, PetId) && strings.Contains(line, "进化") && strings.Contains(line, "JinHua") {
			arrays := strings.Split(strings.Replace(line, ")", "(", -1), "(")
			value := arrays[1]
			array2 := strings.Split(value, ",")
			EvaulateRoute := array2[0]
			ArticlePid := array2[2]
			ArticleNameLine := lines[index-2]
			ArticleName := strings.Split(strings.Replace(ArticleNameLine, ">", "<", -1), "<")[2]
			EvaluateArticle := &model.Article{
				Name: ArticleName,
				Pid:  ArticlePid,
			}
			if EvaulateRoute == "1" {
				EvaulateArticleRoute1 = EvaluateArticle
			}
			if EvaulateRoute == "2" {
				EvaulateArticleRoute2 = EvaluateArticle
			}
		}
	}
	return EvaulateArticleRoute1, EvaulateArticleRoute2
}

func (inst *fusionRepositoryKdhsImpl) Evaluate(mainPet *model.Pet, evaluateRoute string, pid string, bhid string) (bool, error) {
	params := util.InitParam()
	params["n"] = evaluateRoute
	params["id"] = mainPet.Id
	params["bhid"] = bhid
	params["pids"] = pid
	result := CallServerGetInterface("function/jhGate.php", params)

	switch result {
	case "3":
		log.Error("宝宝的等级太低，还不能进化！")
		return false, errors.New("宝宝的等级太低，还不能进化！")
	case "2":
		log.Error("%s %s 缺少进化必须品！", mainPet.Id, mainPet.Name)
		return false, errors.New("缺少进化必须品！")
	case "5":
		log.Error("您没有足够金币进行进化！")
		return false, errors.New("您没有足够金币进行进化！")
	case "6":
		log.Error("%s %s 您已经达到最大的进化次数了！", mainPet.Id, mainPet.Name)
		return false, errors.New("您已经达到最大的进化次数了！")
	case "100":
		log.Error("进化失败 重试")
		return inst.Evaluate(mainPet, evaluateRoute, pid, bhid)
	case "1":
		log.Info("进化成功")
		pet, err := PetRepositoryImpl4KDHS.GetPetDetail(mainPet.Id)
		if err != nil {
			log.Error("获取宠物详情失败")
			return false, errors.New("获取宠物详情失败")
		}
		mainPet.Cc = pet.Cc
		mainPet.Name = pet.Name
		mainPet.Level = pet.Level
		log.Info("当前宠物 %s 成长: %f", pet.Name, mainPet.Cc)
		return true, nil
	default:
		log.Error("%s %s 未知错误 %s", mainPet.Id, mainPet.Name, result)
		return false, errors.New("未知错误")
	}
}

func (inst *fusionRepositoryKdhsImpl) GetPetTypeList() []string {
	//TODO implement me
	panic("implement me")
}

func (inst *fusionRepositoryKdhsImpl) GetExperienceTypeList() []string {
	var ExperienceTypeList = []string{}
	for key := range experienceTypeDict {
		ExperienceTypeList = append(ExperienceTypeList, key)
	}
	return ExperienceTypeList
}

func (inst *fusionRepositoryKdhsImpl) GetExperienceList(ExperienceType string) []string {
	return experienceTypeDict[ExperienceType]
}

func (inst *fusionRepositoryKdhsImpl) GetProtectArticleTypeList() []string {
	var ProtectTypeList = []string{}
	for key := range protectArticleDict {
		ProtectTypeList = append(ProtectTypeList, key)
	}
	return ProtectTypeList
}

func (inst *fusionRepositoryKdhsImpl) GetProjectArticleList(ProtectArticleType string) []string {
	return protectArticleDict[ProtectArticleType]
}

func (inst *fusionRepositoryKdhsImpl) GetNirvanaArticleTypeList() []string {
	var ProtectTypeList = []string{}
	for key := range nirvanaArticleDict {
		ProtectTypeList = append(ProtectTypeList, key)
	}
	return ProtectTypeList
}
func (inst *fusionRepositoryKdhsImpl) GetNirvanaArticleList(ProtectArticleType string) []string {
	return nirvanaArticleDict[ProtectArticleType]
}
