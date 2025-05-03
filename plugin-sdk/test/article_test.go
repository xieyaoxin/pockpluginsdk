package test

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/article"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/user"

	"testing"
)

func TestQueryArticleList(t *testing.T) {
	User := GetLoginUser()
	user.UserServiceInstance.Login(User.LoginName, User.Password)
	articleList, err := article.ArticleServiceInstance.QueryArticleList("")
	if err != nil {
		return
	}
	for _, article := range articleList {
		plugin_log.Info("获取物品列表 %v", article)
	}
}
