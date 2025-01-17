package impl

import (
	"pock_plugins/backend/model"
	"pock_plugins/backend/repository"
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
