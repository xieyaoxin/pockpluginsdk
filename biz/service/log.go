package service

import (
	"pock_plugins/backend/log"
)

func (app *App) GetLog() string {
	return buildSuccessResponse(log.ClearLogList())
}
