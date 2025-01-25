package test

import (
	"fmt"
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"testing"
)

func TestFusion(t *testing.T) {
	GetLoginUser()
	plugin_sdk.InitMergeArticleCache()
	//_, err := plugin_sdk.UserServiceInstance.Login(User.LoginName, User.Password)
	MergeConfig := model.MergeGodConfig{
		MainPet: &model.MergeDragonConfig{
			MainPet: &model.SingleMergeConfig{
				MainPetConfig: model.MergePetConfig{
					PetType: "WX",
					Evaluate: []*model.EvaluateConfig{
						&model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						},
					},
				},
				ProtectType1: "至尊",
				ProtectType2: "3XCC",
			},
			AtePet: &model.SingleMergeConfig{
				MainPetConfig: model.MergePetConfig{
					PetType: "BMW",
				},
				ProtectType1: "3XCC",
				ProtectType2: "3XCC",
			},
			EatPet: model.SingleMergeConfig{
				MainPetConfig: model.MergePetConfig{
					ExperienceType: "1E",
					PetLevel:       60,
					PetCc:          56,
					Evaluate: []*model.EvaluateConfig{
						&model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						}, &model.EvaluateConfig{
							EvaluateRoute: "1",
							ForceEvaluate: true,
						},
					},
				},
				AtePetConfig: model.MergePetConfig{
					ExperienceType: "1E",
					PetLevel:       60,
				},
				ProtectType1: "至尊",
				ProtectType2: "3XCC",
			},
		},
		AteDragon: &model.MergeDragonConfig{
			MainPet: &model.SingleMergeConfig{
				MainPetConfig: model.MergePetConfig{
					PetType:  "LD",
					PetLevel: 60,
				},
				ProtectType1: "至尊",
				ProtectType2: "3XCC",
			},
		},
		EatDragon: &model.SingleMergeConfig{
			MainPetConfig: model.MergePetConfig{
				ExperienceType: "1E",
				PetLevel:       60,
				Evaluate: []*model.EvaluateConfig{
					&model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					}, &model.EvaluateConfig{
						EvaluateRoute: "1",
						ForceEvaluate: true,
					},
				},
			},
			AtePetConfig: model.MergePetConfig{
				ExperienceType: "1E",
				PetLevel:       60,
			},
			ProtectType1: "至尊",
			ProtectType2: "3XCC",
		},
	}
	god, err := plugin_sdk.MergeGod(MergeConfig)
	if err != nil {
		return
	}
	fmt.Printf("合神结果 : %v", god)
}

//
//
//
//
//
//
//	MainPet: &model.SingleMergeConfig{
//		MainPetConfig: model.MergePetConfig{
//			PetType: "WX",
//		},
//		ProtectType1: "至尊",
//		ProtectType2: "3XCC",
//	},
//	AtePet: &model.SingleMergeConfig{
//		MainPetConfig: model.MergePetConfig{
//			PetType: "BMW",
//		},
//		ProtectType1: "3XCC",
//		ProtectType2: "3XCC",
//	},
//	EatPet: model.SingleMergeConfig{
//		MainPetConfig: model.MergePetConfig{
//			ExperienceType: "1E",
//			Evaluate: []*model.EvaluateConfig{
//				&model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true}, &model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true},
//				&model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true}, &model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true},
//				&model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true}, &model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true},
//				&model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true}, &model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true},
//				&model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true}, &model.EvaluateConfig{EvaluateRoute: "2", ForceEvaluate: true},
//			},
//		},
//		AtePetConfig: model.MergePetConfig{
//			ExperienceType: "1E",
//		},
//		ProtectType1: "至尊",
//		ProtectType2: "3XCC",
//	},
//	AteDragon: &model.SingleMergeConfig{
//		MainPetConfig: model.MergePetConfig{
//			PetType: "LD",
//		},
//	},
//	//EatDragon:
//}

//}
