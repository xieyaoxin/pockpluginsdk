package test

import (
	"fmt"
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"

	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
	"testing"
	"time"
)

func TestFightOneTime(t *testing.T) {
	User := GetLoginUser()
	_, err := repository.GetUserRepository().Login(User)
	if err != nil {
		fmt.Printf("登录失败")
		return
	}
	pets, _ := plugin_sdk.PetServiceInstance.GetCarriedPetList()
	pet := pets[0]
	type args struct {
		BattleConfig model.BattleConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				BattleConfig: model.BattleConfig{
					PetId:              pet.Id,
					SkillId:            "1",
					MapId:              "10",
					Difficulty:         "1",
					SkipMonsters:       []string{},
					CatchPets:          []string{},
					RunWhenCatchFailed: true,
					RunWhenNotCatch:    true,
					Balls:              []string{"涅盘兽·大师精灵球（绑定）", "涅盘兽·大师精灵球"},
					Rubbish:            []string{},
					CatchHpThreshold:   100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plugin_sdk.FightOneTime(tt.args.BattleConfig); got != tt.want {
				t.Errorf("FightOneTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFight(t *testing.T) {
	GetLoginUser()

	pets, _ := plugin_sdk.PetServiceInstance.GetCarriedPetList()
	pet := pets[0]
	type args struct {
		BattleConfig model.BattleConfig
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				BattleConfig: model.BattleConfig{
					PetId:              pet.Id,
					SkillId:            "1",
					MapId:              "166",
					Difficulty:         "1",
					SkipMonsters:       []string{},
					CatchPets:          []string{"涅"},
					RunWhenCatchFailed: true,
					RunWhenNotCatch:    true,
					Balls:              []string{"涅盘兽地图·大师精灵球（绑定）"},
					Rubbish:            []string{},
					CatchHpThreshold:   100,
					SaveAfterCatch:     true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plugin_sdk.BattleServiceImpl.Fight(tt.args.BattleConfig, nil)
		})
	}
	plugin_log.Info("当前战斗状态: %v", status.IsBattleRunning())
	time.Sleep(time.Duration(30000000) * time.Second)
	status.SetBattleStatus(status.Waiting2Stop)
	plugin_log.Info("当前战斗状态: %v", status.FightStatus.BattleStatus)
	time.Sleep(time.Duration(20) * time.Second)
	plugin_log.Info("当前战斗状态: %v", status.FightStatus.BattleStatus)

}
