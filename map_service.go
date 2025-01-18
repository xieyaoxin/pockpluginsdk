package plugin_sdk

import (
	"github.com/xieyaoxin/plugin-sdk/biz/model"
	"github.com/xieyaoxin/plugin-sdk/biz/repository"
)

var MapServiceInstance = &mapService{}

type mapService struct {
}

func (*mapService) GetBattleMapList() []*model.BattleMap {
	return repository.GetMapRepository().GetMapList()
}
