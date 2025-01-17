package repository

import "pock_plugins/backend/model"

type UserRepository interface {
	Login(user model.User) (*model.User, error)
	GetUserInfo() (*model.User, error)
}
