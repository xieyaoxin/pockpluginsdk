package plugin_sdk

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
)

var UserServiceInstance = &userService{}

type userService struct {
}

func (*userService) Login(loginName, password string) (*model.User, error) {
	user := model.User{
		LoginName: loginName,
		Password:  password,
	}
	TempToken := util.MD5(user.LoginName + util.GenerateRandomSeed())
	user.TempToken = TempToken
	status.SetLoginUser(user)
	return repository.GetUserRepository().Login(user)
}
