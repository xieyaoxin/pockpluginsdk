package repository

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
)

type UserRepository interface {
	Login(user model.User) (*model.User, error)
	GetUserInfo() (*model.User, error)
}
