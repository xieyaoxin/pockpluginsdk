package impl

import (
	"plugin-sdk/biz/model"
	"plugin-sdk/biz/repository"
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
