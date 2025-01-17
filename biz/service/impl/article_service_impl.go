package impl

import (
	"pock_plugins/backend/model"
	"pock_plugins/backend/repository"
)

var ArticleServiceInstance = &articleService{}

type articleService struct {
}

func (*articleService) QueryArticleList(name string) ([]*model.Article, error) {
	return repository.GetArticleRepository().GetArticles(name)
}
