package repository

import "github.com/xieyaoxin/plugin-sdk/biz/model"

type UserRepository interface {
	Login(user model.User) (*model.User, error)
	GetUserInfo() (*model.User, error)
}
