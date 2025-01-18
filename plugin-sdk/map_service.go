package plugin_sdk

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
)

var MapServiceInstance = &mapService{}

type mapService struct {
}

func (*mapService) GetBattleMapList() []*model.BattleMap {
	return repository.GetMapRepository().GetMapList()
}
