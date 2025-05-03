package test

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/equip"

	"testing"
)

func TestGetEquips(t *testing.T) {
	GetLoginUser()
	equips := plugin_sdk.EquipServiceImplInstance.GetEquip("8178204")
	for _, equip := range equips {
		plugin_log.Info("装备名称: %s 装备ID %s", equip.Name, equip.EquipId)
	}
}

func TestOffEquips(t *testing.T) {
	GetLoginUser()
	plugin_sdk.EquipServiceImplInstance.OffEquip("8178204")
}
