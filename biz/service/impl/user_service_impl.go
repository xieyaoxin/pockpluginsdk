package impl

import (
	"plugin-sdk/backend/model"
	"plugin-sdk/backend/repository"
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
