package test

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"testing"
)

func TestGetFarmedPet(t *testing.T) {
	GetLoginUser()
	for {
		pets, _ := repository.GetPetRepository().GetCarriedPets()
		for _, pet := range pets {
			plugin_log.Info("名称 %s  ", pet.Name)
		}
		plugin_log.Info("")
		plugin_log.Info("")

	}
}
