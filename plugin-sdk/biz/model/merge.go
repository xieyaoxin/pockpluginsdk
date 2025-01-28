package model

type MergeGodConfig struct {
	MainPet   *MergeDragonConfig // 初始化主宠
	AteDragon *MergeDragonConfig // 副龙配置
	EatDragon *SingleMergeConfig // 副龙配置
}

type MergeDragonConfig struct {
	MainPet *SingleMergeConfig // 初始化主宠
	AtePet  *SingleMergeConfig // 合成副宠
	EatPet  SingleMergeConfig  // 吃副宠的配置
}
type SingleMergeConfig struct {
	MainPetConfig MergePetConfig
	AtePetConfig  MergePetConfig
	ProtectType1  string
	ProtectType2  string
}

type MergePetConfig struct {
	PetType        string
	PetId          string
	ExperienceType string
	PetLevel       int64
	PetCc          float64
	Evaluate       []*EvaluateConfig
}

type EvaluateConfig struct {
	EvaluateRoute string `json:"evaluateRoute"`
	ForceEvaluate bool   `json:"forceEvaluate"`
}

type NirvanaPetConfig struct {
	MergePetConfig
	PetName   string
	UseEgg    bool
	IsNirvana bool
}
type NirvanaConfig struct {
	MainPet      NirvanaPetConfig
	AtePet       NirvanaPetConfig
	NirvanaPet   NirvanaPetConfig
	ProtectType1 string
	ProtectType2 string
}

// 宠物类型
const (
	WX  = "WX"  // 五系
	LD  = "LD"  // 龙蛋
	ALY = "ALY" // 爱丽哑
	BMW = "BMW" // 波姆王
	BM  = "BM"  // 波姆
)

var petTypeMap = make(map[string][]string)

func CopySingleMergeConfig(config SingleMergeConfig) SingleMergeConfig {
	return SingleMergeConfig{
		MainPetConfig: CopyMergePetConfig(config.MainPetConfig),
		AtePetConfig:  CopyMergePetConfig(config.AtePetConfig),
		ProtectType1:  config.ProtectType1,
		ProtectType2:  config.ProtectType2,
	}
}

func CopyMergePetConfig(config MergePetConfig) MergePetConfig {
	return MergePetConfig{
		PetType:        config.PetType,
		PetId:          config.PetId,
		ExperienceType: config.ExperienceType,
		PetLevel:       config.PetLevel,
		PetCc:          config.PetCc,
		Evaluate:       CopyEvaluates(config.Evaluate),
	}
}

func CopyEvaluates(evaluate []*EvaluateConfig) []*EvaluateConfig {
	Temp := make([]*EvaluateConfig, len(evaluate))
	copy(Temp, evaluate)
	return Temp
}

const (
	EXPERIENCE_200W  = "200W"  // 200W月饼
	EXPERIENCE_500W  = "500W"  // 500W月饼
	EXPERIENCE_5000W = "5000W" // 5000W月饼
	EXPERIENCE_1E    = "1E"    // 册子
	EXPERIENCE_10E   = "10E"   // 10E
	EXPERIENCE_15E   = "15E"   // 15E
	EXPERIENCE_20E   = "20E"   // 20E
	EXPERIENCE_200E  = "200E"  // 20E
	EXPERIENCE_800E  = "800E"  // 10E
)

// 保护石头1
const (
	PROTECT_HC   = "护宠"   // 护宠
	PROTECT_SH   = "守魂"   // 守混
	PROTECT_TS   = "天神"   // 天神
	PROTECT_ZZ   = "至尊"   // 天神
	PROTECT_3XCC = "3XCC" // 3*
	PROTECT_3X   = "3X"   // 3*

	PROTECT_NIRVANA_SD  = "神丹" // 神丹
	PROTECT_NIRVANA_NPD = "涅丹" // 神丹
	PROTECT_NIRVANA_SSD = "圣丹" // 神丹

	PROTECT_NIRVANA_YPNC = "一品捏成" // 神丹
	PROTECT_NIRVANA_SPNC = "上品捏成" // 神丹
	PROTECT_NIRVANA_JPNC = "极品捏成" // 神丹
	PROTECT_NIRVANA_SPMZ = "上品命中" // 神丹
	PROTECT_NIRVANA_JPMZ = "极品命中" // 神丹
	PROTECT_NIRVANA_JPNG = "极品涅攻" // 神丹
)

var NirvanaEggList = []string{
	"涅槃兽之卵(限时)", "涅槃兽之卵(绑定)", "涅盘兽之卵",
}

var DragonEggList = []string{
	"小神龙琅玡之卵(限时)", "小神龙琅玡之卵(绑)", "小神龙琅玡之卵",
}
