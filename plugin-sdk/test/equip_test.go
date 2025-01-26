package test

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"testing"
)

func TestGetEquips(t *testing.T) {
	GetLoginUser()
	equips := plugin_sdk.EquipServiceImplInstance.GetEquip("8178204")
	for _, equip := range equips {
		log.Info("装备名称: %s 装备ID %s", equip.Name, equip.EquipId)
	}
}

func TestOffEquips(t *testing.T) {
	GetLoginUser()
	plugin_sdk.EquipServiceImplInstance.OffEquip("8178204")
}
