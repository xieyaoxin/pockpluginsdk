package plugin_sdk

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"time"
)

var TimerConfig = []*model.TimerConfig{
	// 自动清理垃圾
	&model.TimerConfig{
		Enable:       true,
		Schedule:     5,
		TimeTaskType: "自动清理物品",
		Config: model.DropArticleTimerConfig{
			DropArticleList:      []string{},
			DropArticleBlackList: []string{"★★★", "卵", "礼包", "天神", "00w经验月饼", "水晶卡", "亿经验卷", "修炼仙册", "至尊神石", "强化丹", "进化宝石", "天仙", "玉露", "保底石", "扫雷", "玛雅", "女神", "九星龙珠", "八星龙珠", "七星龙珠", "六星龙珠", "涅盘兽", "涅槃兽", "涅盘神丹", "涅盘圣丹", "宝宝祭石", "贤者之石", "女神圣水", "圣诞", "婵娟", "金乌", "后羿", "嫦娥", "真武神印", "梦想龙珠", "神圣攻击珠", "神圣命中珠", "江雪之翼", "升级", "守护印记", "k歌之王", "辉煌", "元宝"},
		},
	},
	// 自动副本
	&model.TimerConfig{
		Enable:       true,
		Schedule:     10,
		TimeTaskType: "自动副本",
		Config: model.DropArticleTimerConfig{
			DropArticleList:      []string{},
			DropArticleBlackList: []string{"★★★", "卵", "礼包", "天神", "00w经验月饼", "水晶卡", "亿经验卷", "修炼仙册", "至尊神石", "强化丹", "进化宝石", "天仙", "玉露", "保底石", "扫雷", "玛雅", "女神", "九星龙珠", "八星龙珠", "七星龙珠", "六星龙珠", "涅盘兽", "涅槃兽", "涅盘神丹", "涅盘圣丹", "宝宝祭石", "贤者之石", "女神圣水", "圣诞", "婵娟", "金乌", "后羿", "嫦娥", "真武神印", "梦想龙珠", "神圣攻击珠", "神圣命中珠", "江雪之翼", "升级", "守护印记", "k歌之王", "辉煌", "元宝"},
		},
	},
}

func init() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	var Times int64
	// 启动一个 goroutine 来处理定时任务
	go func() {
		for {
			select {
			case <-ticker.C:
				handleTimer(Times)
				Times++
			}
		}
	}()
}

func handleTimer(Times int64) {
	// 丢弃物品
	// 自动空投
	// 自动副本
}

func autoDungeonInstance() {
	//	 判断和记录当前任务类型和状态
	//	 合神跳过
	//	 其他的将任务状态设置为暂停

	//	 接收
}
