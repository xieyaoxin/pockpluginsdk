package plugin_sdk

import (
	"github.com/xieyaoxin/plugin-sdk/biz/model"
	"github.com/xieyaoxin/plugin-sdk/biz/repository"
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
