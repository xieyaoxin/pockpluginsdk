package timer

import (
	"encoding/json"
	"errors"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/chain/config"
	"time"
)

var TimerConfigMap = make(map[string]*model.TimerConfig)
var TimerHandlerMap = make(map[string]model.TimerHandleInterface)

//var TimerConfig = []*model.TimerConfig{
//	// 自动清理垃圾
//	&model.TimerConfig{
//		Enable:       true,
//		Schedule:     5,
//		TimeTaskType: "自动清理物品",
//		Config: model.DropArticleTimerConfig{
//			DropArticleList:      []string{},
//			DropArticleBlackList: []string{"★★★", "卵", "礼包", "天神", "00w经验月饼", "水晶卡", "亿经验卷", "修炼仙册", "至尊神石", "强化丹", "进化宝石", "天仙", "玉露", "保底石", "扫雷", "玛雅", "女神", "九星龙珠", "八星龙珠", "七星龙珠", "六星龙珠", "涅盘兽", "涅槃兽", "涅盘神丹", "涅盘圣丹", "宝宝祭石", "贤者之石", "女神圣水", "圣诞", "婵娟", "金乌", "后羿", "嫦娥", "真武神印", "梦想龙珠", "神圣攻击珠", "神圣命中珠", "江雪之翼", "升级", "守护印记", "k歌之王", "辉煌", "元宝"},
//		},
//	},
//	// 自动副本

//}

func init() {
	TimerHandlerMap["自动清理物品"] = &DropArticleTimerHandler{}
	TimerHandlerMap["自动副本"] = &DungeonInstanceTimerHandler{}

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		var Times int64
		// 启动一个 goroutine 来处理定时任务
		func() {
			for {
				select {
				case tick := <-ticker.C:
					plugin_log.Info("%v", tick)
					Times++
					handleTimer(Times)
				}
			}
		}()
	}()
}

func handleTimer(Times int64) {
	plugin_log.Info("第 %d 次定时循环", Times)
	TimerTasks := []*model.TimerConfig{}
	for _, v := range TimerConfigMap {
		if v.Enable && v.Schedule > 0 && Times%v.Schedule == 0 {
			TimerTasks = append(TimerTasks, v)
		}
	}

	for _, TimerConfig := range TimerTasks {
		if TimerConfig.Handle == nil {
			plugin_log.Error("定时任务%s，未配置定时任务回调", TimerConfig.TimeTaskType)
			continue
		}
		TimerConfig.Handle.HandleTimer(TimerConfig.Config)
	}
}

func InitTimer(Account string) {
	ConfigList := config.GetDefaultTimerConfig(Account)
	for _, Config := range ConfigList {
		Handle := TimerHandlerMap[Config.TimeTaskType]
		Config.Handle = Handle
		UpdateTimer(*Config)
	}
}

// 修改定时任务
func UpdateTimer(config model.TimerConfig) error {
	configStr, _ := json.Marshal(config)
	plugin_log.Info("定时配置： %s", string(configStr))
	if OriginConfig, exists := TimerConfigMap[config.TimeTaskType]; exists {
		OriginConfig.Config = config.Config
		OriginConfig.Enable = config.Enable
		OriginConfig.Schedule = config.Schedule
		OriginConfig.Handle = config.Handle
	} else {
		plugin_log.Error("无法找到定时任务  %s", config.TimeTaskType)
		return errors.New("无法找到定时任务  " + config.TimeTaskType)
	}
	return nil
}
