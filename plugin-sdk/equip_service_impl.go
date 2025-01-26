package plugin_sdk

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
)

var EquipServiceImplInstance = &equipServiceImpl{}

type equipServiceImpl struct {
}

func (*equipServiceImpl) GetEquip(PetId string) []*model.Equip {
	return repository.GetEquipRepository().GetEquipList(PetId)
}

func (inst *equipServiceImpl) OffEquip(PetId string) {
	// 获取身上装备
	equips := inst.GetEquip(PetId)
	for _, equip := range equips {
		result := repository.GetEquipRepository().OffEquip(PetId, equip.EquipPId)
		log.Info("脱装备： %s , 结果 %b", equip.Name, result)
	}
}
