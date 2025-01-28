package test

import (
	"fmt"
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"io"
	"os"
	"strings"
)

func init() {
	// 初始化区服
	status.SERVER_NAME = status.KDHS
}

func GetLoginUser() model.User {

	f, err := os.Open("config.txt")
	if err != nil {
		// 打开文件失败
		log.Fatal(err.Error())
	}
	var data []byte
	buf := make([]byte, 1024)
	for {
		// 将文件中读取的byte存储到buf中
		n, err1 := f.Read(buf)
		if err1 != nil && err != io.EOF {
			log.Fatal(err1.Error())
		}
		if n == 0 {
			break
		}
		// 将读取到的结果追加到data切片中
		data = append(data, buf[:n]...)
	}
	fmt.Printf(string(data))
	properties := strings.Split(string(data), "\n")
	userInfo := properties[0]
	userInfoArray := strings.Split(userInfo, ";")
	User := model.User{
		LoginName: userInfoArray[0],
		Password:  userInfoArray[1],
	}
	login, err := plugin_sdk.UserServiceInstance.Login(User.LoginName, User.Password)
	if err != nil {
		return model.User{}
	}
	log.Info("登录成功 %v", login)
	return User
}

func HoldOn() {
	ch := make(chan struct{})
	// 使用select语句阻塞主协程
	select {
	case <-ch:
	}
}
