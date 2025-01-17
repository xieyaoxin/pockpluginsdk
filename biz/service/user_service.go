package service

import (
	"fmt"
	"pock_plugins/backend/service/impl"
)

func (app *App) GetUserInfo() string {
	return ""
}

func (app *App) Login2Server(LoginName, Password string) string {
	fmt.Printf("userName %s   Password %s", LoginName, Password)
	response, err := impl.UserServiceInstance.Login(LoginName, Password)
	if err != nil {
		return buildFailedResponse(response)
	}
	return buildSuccessResponse(response)
}
