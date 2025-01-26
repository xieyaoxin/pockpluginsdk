package test

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"testing"
)

func TestGetFarmedPet(t *testing.T) {
	GetLoginUser()
	repository.GetPetRepository().GetFarmedPets()

}
