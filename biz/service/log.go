package service

import (
	"plugin-sdk/biz/log"
)

func (app *App) GetLog() string {
	return buildSuccessResponse(log.ClearLogList())
}
