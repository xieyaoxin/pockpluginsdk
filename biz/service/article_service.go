package service

import (
	"pock_plugins/backend/log"
	"pock_plugins/backend/service/impl"
)

func (app *App) GetArticleList(name string) string {
	articles, err := impl.ArticleServiceInstance.QueryArticleList(name)
	log.Info("查询物品: ", name)
	if err != nil {
		return buildFailedResponse(articles)
	}
	return buildSuccessResponse(articles)
}
