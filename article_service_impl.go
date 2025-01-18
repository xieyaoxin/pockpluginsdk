package plugin_sdk

import (
	"github.com/xieyaoxin/plugin-sdk/biz/model"
	"github.com/xieyaoxin/plugin-sdk/biz/repository"
)

var ArticleServiceInstance = &articleService{}

type articleService struct {
}

func (*articleService) QueryArticleList(name string) ([]*model.Article, error) {
	return repository.GetArticleRepository().GetArticles(name)
}
