package plugin_sdk

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
)

var UserServiceInstance = &userService{}

type userService struct {
}

func (*userService) Login(loginName, password string) (*model.User, error) {
	return repository.GetUserRepository().Login(model.User{
		LoginName: loginName,
		Password:  password,
	})
}
