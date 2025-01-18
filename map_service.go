package plugin_sdk

import (
	"plugin-sdk/biz/model"
	"plugin-sdk/biz/repository"
)

var MapServiceInstance = &mapService{}

type mapService struct {
}

func (*mapService) GetBattleMapList() []*model.BattleMap {
	return repository.GetMapRepository().GetMapList()
}
