package repository

import "pock_plugins/backend/model"

type MapRepository interface {
	GetMapList() []*model.BattleMap
}
