package test

import (
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	chains "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/timer"
	"testing"
)

//8206429
//

func TestTtFight(t *testing.T) {
	GetLoginUser()
	chains.TtChain()

	plugin_sdk.UpdateTimer(model.TimerConfig{
		Enable:       true,
		Schedule:     10,
		TimeTaskType: "自动清理物品",
		Config: model.DropArticleTimerConfig{
			DropArticleList: []string{
				"追龙令", "龙魂", "黑魔杖", "黑魔珠", "黑魔盔", "黑魔铠", "黑魔靴",
			},
			DropArticleBlackList: []string{"★★★", "龙魂礼包", "卵", "礼包", "天神", "00w经验月饼", "水晶卡", "亿经验卷", "修炼仙册", "至尊神石", "强化丹", "进化宝石", "天仙", "玉露", "保底石", "扫雷", "玛雅", "女神", "九星龙珠", "八星龙珠", "七星龙珠", "六星龙珠", "涅盘兽", "涅槃兽", "涅盘神丹", "涅盘圣丹", "宝宝祭石", "贤者之石", "女神圣水", "圣诞", "婵娟", "金乌", "后羿", "嫦娥", "真武神印", "梦想龙珠", "神圣攻击珠", "神圣命中珠", "江雪之翼", "升级", "守护印记", "k歌之王", "辉煌", "元宝"},
		},
		Handle: &timer.DropArticleTimerHandler{},
	})

	plugin_sdk.UpdateTimer(model.TimerConfig{
		Enable:       true,
		Schedule:     30,
		TimeTaskType: "自动副本",
		Config: model.DungeonInstanceTimerConfig{
			PetId:     "",
			PetName:   "辣椒",
			SkillId:   "799",
			SkillName: "",
			MapList:   []string{"11", "12", "144"},
		},
		Handle: &timer.DungeonInstanceTimerHandler{},
	})
	HoldOn()
}
