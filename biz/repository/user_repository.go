package repository

import "plugin-sdk/biz/model"

type UserRepository interface {
	Login(user model.User) (*model.User, error)
	GetUserInfo() (*model.User, error)
}
