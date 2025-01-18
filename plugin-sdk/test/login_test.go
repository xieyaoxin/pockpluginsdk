package test

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"testing"
)

func TestLogin(t *testing.T) {
	User := GetLoginUser()
	login, err := repository.GetUserRepository().Login(User)
	if err != nil {
		return
	}
	log.Info("登陆成功 %v", login)
}
