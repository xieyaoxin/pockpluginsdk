package test

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"

	"testing"
)

func TestQueryArticleList(t *testing.T) {
	User := GetLoginUser()
	plugin_sdk.UserServiceInstance.Login(User.LoginName, User.Password)
	articleList, err := plugin_sdk.ArticleServiceInstance.QueryArticleList("")
	if err != nil {
		return
	}
	for _, article := range articleList {
		plugin_log.Info("获取物品列表 %v", article)

	}
}
