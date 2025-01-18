package cqtt

import (
	"github.com/xieyaoxin/plugin-sdk/plugin-sdk/biz/model"
)

var MapRepositoryImpl4CQTTInstance = &mapRepositoryImpl4CQTT{}

type (
	mapRepositoryImpl4CQTT struct {
	}
)

func (*mapRepositoryImpl4CQTT) GetMapList() []*model.BattleMap {
	// 基础地图
	BaseMapList := []*model.BattleMap{
		{
			MapId:   "1",
			MapName: "新手训练营",
		}, {
			MapId:   "2",
			MapName: "妖精森林",
		}, {
			MapId:   "3",
			MapName: "潮汐海涯",
		}, {
			MapId:   "4",
			MapName: "巨石山脉",
		}, {
			MapId:   "5",
			MapName: "黄金陵",
		}, {
			MapId:   "6",
			MapName: "炽热沙滩",
		}, {
			MapId:   "7",
			MapName: "尤玛火山",
		}, {
			MapId:   "8",
			MapName: "死亡沙漠",
		}, {
			MapId:   "9",
			MapName: "海市盛楼",
		}, {
			MapId:   "10",
			MapName: "冰滩",
		}, {
			MapId:   "15",
			MapName: "圣诞小屋",
		}, {
			MapId:          "16",
			MapName:        "海底世界",
			DifficultyFlag: true,
		},
	}
	NewLandMapList := []*model.BattleMap{
		{
			MapId:          "100",
			MapName:        "石阵",
			DifficultyFlag: true,
		}, {
			MapId:          "103",
			MapName:        "平原",
			DifficultyFlag: true,
		}, {
			MapId:          "106",
			MapName:        "绿荫林",
			DifficultyFlag: true,
		}, {
			MapId:          "109",
			MapName:        "五指石印",
			DifficultyFlag: true,
		}, {
			MapId:          "112",
			MapName:        "鬼屋",
			DifficultyFlag: true,
		}, {
			MapId:          "115",
			MapName:        "天空之城",
			DifficultyFlag: true,
		}, {
			MapId:          "118",
			MapName:        "天之路",
			DifficultyFlag: true,
		}, {
			MapId:          "121",
			MapName:        "危之路",
			DifficultyFlag: true,
		}, {
			MapId:          "100",
			MapName:        "石阵",
			DifficultyFlag: true,
		},
	}
	return append(BaseMapList, NewLandMapList...)
}
