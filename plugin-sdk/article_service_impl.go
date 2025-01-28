package plugin_sdk

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_sdk_const"
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

func (*articleService) UseSjk() error {
	//	 先开水晶卡
	log.Info("身上水晶不够,查询身上水晶卡信息")
	ArticleList, err := ArticleServiceInstance.QueryArticleListByNameLists(plugin_sdk_const.SJK_LIST_NAME)
	if err != nil || ArticleList == nil || len(ArticleList) == 0 {
		return errors.New("获取水晶卡失败")
	}
	log.Info("使用水晶卡 %s", ArticleList[0].Name)
	err = ArticleServiceInstance.UserArticle(ArticleList[0])
	if err != nil {
		return err
	}
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
