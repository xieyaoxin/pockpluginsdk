package service

import "plugin-sdk/biz/service/impl"

func (app *App) GetMapList() string {
	return buildSuccessResponse(impl.MapServiceInstance.GetBattleMapList())
}
