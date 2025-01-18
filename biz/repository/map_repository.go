package repository

import "github.com/xieyaoxin/plugin-sdk/biz/model"

type MapRepository interface {
	GetMapList() []*model.BattleMap
}
