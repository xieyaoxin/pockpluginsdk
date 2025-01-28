package test

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"

	"testing"
)

func TestLogin(t *testing.T) {
	User := GetLoginUser()
	login, err := plugin_sdk.UserServiceInstance.Login(User.LoginName, User.Password)
	if err != nil {
		return
	}
	plugin_log.Info("登陆成功 %v", login)
}
