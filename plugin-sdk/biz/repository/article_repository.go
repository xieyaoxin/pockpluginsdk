package repository

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
)

type ArticleRepository interface {
	UseArticle(articleId string) error
	GetArticles(name string) ([]*model.Article, error)
	GetArticleDetail(articleId string) model.Article
}
