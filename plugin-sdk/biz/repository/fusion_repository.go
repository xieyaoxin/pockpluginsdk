package repository

import "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"

type FusionRepository interface {
	// Fusion 合成
	Fusion(Pet1 model.Pet, Pet2 model.Pet, protect1 model.Article, article2 model.Article) (bool, error)
	// Nirvana 涅槃
	Nirvana(Pet1 model.Pet, Pet2 model.Pet, NPS model.Pet, protect1 model.Article, article2 model.Article) (bool, error)
	// Evaluate 进化
	Evaluate(mainPet *model.Pet, evaluateRoute string, pid string, bhid string) (bool, error)

	GetEvaluateArticle(PetId string) (*model.Article, *model.Article)
	GetPetTypeList() []string
	GetExperienceTypeList() []string
	GetExperienceList(ExperienceType string) []string
	GetProtectArticleTypeList() []string
	GetProjectArticleList(ProtectArticleType string) []string

	GetNirvanaArticleTypeList() []string
	GetNirvanaArticleList(ProtectArticleType string) []string
}
