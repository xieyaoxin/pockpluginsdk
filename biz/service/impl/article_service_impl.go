package impl

import (
	"plugin-sdk/biz/model"
	"plugin-sdk/biz/repository"
)

var ArticleServiceInstance = &articleService{}

type articleService struct {
}

func (*articleService) QueryArticleList(name string) ([]*model.Article, error) {
	return repository.GetArticleRepository().GetArticles(name)
}
