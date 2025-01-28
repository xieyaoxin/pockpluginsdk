package kdhs

import (
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"

	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
)

var UserRepositoryImpl4KDHSInstance = &userRepositoryImpl4KDHS{}

type (
	userRepositoryImpl4KDHS struct {
	}
)

func (*userRepositoryImpl4KDHS) Login(user model.User) (*model.User, error) {
	token := Login(user, 0)
	if token == "" {
		plugin_log.Error("登录失败")
		return nil, errors.New("登录失败")
	}
	status.SetLoginUser(model.User{
		LoginName: user.LoginName,
		UserName:  user.UserName,
		Password:  user.Password,
		Token:     token,
	})
	return status.GetLoginUser(), nil
}

func (*userRepositoryImpl4KDHS) GetUserInfo() (*model.User, error) {
	return nil, nil
}
