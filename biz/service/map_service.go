package service

import "pock_plugins/backend/service/impl"

func (app *App) GetMapList() string {
	return buildSuccessResponse(impl.MapServiceInstance.GetBattleMapList())
}
