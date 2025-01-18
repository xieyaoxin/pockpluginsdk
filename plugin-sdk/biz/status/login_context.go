package status

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
)

// 登录态
var currentLoginUser = &model.User{}

func GetLoginUser() *model.User {
	return currentLoginUser
}

func GetLoginToken() string {
	if currentLoginUser.Token == "" {
		return currentLoginUser.TempToken
	}
	return currentLoginUser.Token
}

func SetLoginUser(user model.User) {
	currentLoginUser = &user
}
