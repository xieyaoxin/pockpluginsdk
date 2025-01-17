package cqtt

import (
	"errors"
	"pock_plugins/backend/log"
	"pock_plugins/backend/model"
	"pock_plugins/backend/status"
)

var UserRepositoryImpl4CQTTInstance = &userRepositoryImpl4CQTT{}

type (
	userRepositoryImpl4CQTT struct {
	}
)

func (*userRepositoryImpl4CQTT) Login(user model.User) (*model.User, error) {
	TempToken := MD5(user.LoginName + generateRandomSeed())
	user.TempToken = TempToken
	status.SetLoginUser(user)
	token := Login(user, 0)
	if token == "" {
		log.Error("登录失败")
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

func (*userRepositoryImpl4CQTT) GetUserInfo() (*model.User, error) {
	return nil, nil
}
