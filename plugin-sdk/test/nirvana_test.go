package test

import (
	"encoding/json"
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"

	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"testing"
)

func TestNirvana(t *testing.T) {
	GetLoginUser()
	Config := &model.NirvanaConfig{
		MainPet: model.NirvanaPetConfig{
			MergePetConfig: model.MergePetConfig{
				ExperienceType: "1E",
				PetLevel:       60,
				PetCc:          0,
				Evaluate: []*model.EvaluateConfig{
					&model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "2",
						ForceEvaluate: true,
					},
				},
			},
			PetName: "辣椒",
			UseEgg:  false,
		},
		AtePet: model.NirvanaPetConfig{
			MergePetConfig: model.MergePetConfig{
				PetLevel:       80,
				ExperienceType: "20E",
			},
			PetName: "小神龙",
			UseEgg:  true,
		},
		NirvanaPet: model.NirvanaPetConfig{
			MergePetConfig: model.MergePetConfig{
				PetLevel:       60,
				ExperienceType: "1E",
			},
			PetName:   "涅磐兽",
			UseEgg:    true,
			IsNirvana: true,
		},
		ProtectType1: "神丹",
		ProtectType2: "上品捏成",
	}
	marshal, err2 := json.Marshal(Config)
	if err2 != nil {
		return
	}
	plugin_log.Info("涅槃配置为 %s", string(marshal))
	plugin_sdk.InitNirvanaCache()
	Count := 0
	for Count < 50 {
		_, err := plugin_sdk.NirvanaServiceImplInstance.Nirvana(Config)
		if err != nil {
			return
		}
		Count++
	}

}
