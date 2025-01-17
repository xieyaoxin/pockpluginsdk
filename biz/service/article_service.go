package service

import (
	"plugin-sdk/biz/log"
	"plugin-sdk/biz/service/impl"
)

func (app *App) GetArticleList(name string) string {
	articles, err := impl.ArticleServiceInstance.QueryArticleList(name)
	log.Info("查询物品: ", name)
	if err != nil {
		return buildFailedResponse(articles)
	}
	return buildSuccessResponse(articles)
}
