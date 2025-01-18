package test

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"testing"
)

func TestLogin(t *testing.T) {
	User := GetLoginUser()
	login, err := plugin_sdk.UserServiceInstance.Login(User.LoginName, User.Password)
	if err != nil {
		return
	}
	log.Info("登陆成功 %v", login)
}
