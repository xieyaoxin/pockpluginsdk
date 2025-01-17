package repository

import "plugin-sdk/biz/model"

type MapRepository interface {
	GetMapList() []*model.BattleMap
}
