package test

import (
	"fmt"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/pet"
	"testing"
)

func TestFightOneTime(t *testing.T) {
	User := GetLoginUser()
	_, err := repository.GetUserRepository().Login(User)
	if err != nil {
		fmt.Printf("登录失败")
		return
	}
	pets, _ := pet.PetServiceInstance.GetCarriedPetList()
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
			//if got := chain.BattleServiceImplInstance.FightOneTime(tt.args.BattleConfig); got != tt.want {
			//	t.Errorf("FightOneTime() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestFight(t *testing.T) {
	GetLoginUser()

	pets, _ := pet.PetServiceInstance.GetCarriedPetList()
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
			//chain.BattleServiceImplInstance.FightByConfig(tt.args.BattleConfig, nil)
		})
	}
	//plugin_log.Info("当前战斗状态: %v", status.IsBattleRunning())
	//time.Sleep(time.Duration(30000000) * time.Second)
	//status.SetBattleStatus(status.Waiting2Stop)
	//plugin_log.Info("当前战斗状态: %v", status.GetBattleStatus())
	//time.Sleep(time.Duration(20) * time.Second)
	//plugin_log.Info("当前战斗状态: %v", status.GetBattleStatus())

}

func TestFightXDL(t *testing.T) {
	GetLoginUser()

	pets, _ := pet.PetServiceInstance.GetCarriedPetList()
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
					SkillId:            "743",
					MapId:              "100",
					Difficulty:         "3",
					SkipMonsters:       []string{},
					CatchPets:          []string{},
					RunWhenCatchFailed: false,
					RunWhenNotCatch:    false,
					Balls:              []string{},
					Rubbish:            []string{},
					CatchHpThreshold:   100,
					SaveAfterCatch:     true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//chain.BattleServiceImplInstance.FightByConfig(tt.args.BattleConfig, nil)
			//marshal, err := json.Marshal(tt.args.BattleConfig)
			//if err != nil {
			//	return
			//}
			//plugin_log.Info(string(marshal))
			//	195545
			//	275645
		})
	}
	//plugin_log.Info("当前战斗状态: %v", status.IsBattleRunning())
	//time.Sleep(time.Duration(30000000) * time.Second)
	//status.SetBattleStatus(status.Waiting2Stop)
	//plugin_log.Info("当前战斗状态: %v", status.GetBattleStatus())
	//time.Sleep(time.Duration(20) * time.Second)
	//plugin_log.Info("当前战斗状态: %v", status.GetBattleStatus())

}
