package repository

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
)

type MapRepository interface {
	GetMapList() []*model.BattleMap
}
