package plugin_sdk

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
)

var ArticleServiceInstance = &articleService{}

type articleService struct {
}

func (*articleService) QueryArticleList(name string) ([]*model.Article, error) {
	return repository.GetArticleRepository().GetArticles(name)
}

func (*articleService) QueryArticleListByNameLists(articleNameList []string) ([]*model.Article, error) {
	bagArticles, err := repository.GetArticleRepository().GetArticles("")
	if err != nil {
		return nil, err
	}
	articleList := []*model.Article{}
	for _, articleName := range articleNameList {
		article := getArticleByName(articleName, bagArticles)
		if article == nil {
			continue
		}
		articleList = append(articleList, article)
	}
	return articleList, nil
}

func (*articleService) UserArticle(article *model.Article) error {
	err := repository.GetArticleRepository().UseArticle(article.ID)
	if err != nil {
		return err
	}
	article.ArticleCount = article.ArticleCount - 1
	return nil
}

func getArticleByName(articleName string, articleList []*model.Article) *model.Article {
	for _, article := range articleList {
		if article.Name == articleName {
			return article
		}
	}
	return nil
}
