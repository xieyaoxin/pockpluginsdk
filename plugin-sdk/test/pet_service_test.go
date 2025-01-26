package test

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"testing"
)

func TestGetFarmedPet(t *testing.T) {
	GetLoginUser()
	for {
		pets, _ := repository.GetPetRepository().GetCarriedPets()
		for _, pet := range pets {
			log.Info("名称 %s  ", pet.Name)
		}
		log.Info("")
		log.Info("")

	}
}
