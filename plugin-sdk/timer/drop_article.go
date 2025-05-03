package timer

import (
	"encoding/json"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/plugin_log"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	plugin_sdk "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/operation/article"
)

type DropArticleTimerHandler struct {
}

func (*DropArticleTimerHandler) HandleTimer(OriginConfig interface{}) {

	Config, err := Convert2DropArticleTimerConfig(OriginConfig)
	if err != nil {
		plugin_log.Error("解析定时任务配置失败,%d", OriginConfig)
		return
	}

	Articls, err1 := plugin_sdk.ArticleServiceInstance.QueryArticleList("")
	if err1 != nil {
		plugin_log.Error("获取物品列表失败")
		return
	}
	if Config.DropArticleBlackList == nil || len(Config.DropArticleList) == 0 {
		return
	}
	if Config.DropArticleBlackList == nil {
		Config.DropArticleBlackList = []string{}
	}
	for _, Article := range Articls {
		DropFlag := util.SlicesLikeString(Config.DropArticleList, Article.Name)
		if DropFlag {
			if Article.ArticleType == "礼包类" || Article.ArticleType == "宝箱类" {
				plugin_log.Info("物品: %s 为礼包类 跳过丢弃", Article.Name)
				continue
			}
			BlackFlag := util.SlicesLikeString(Config.DropArticleBlackList, Article.Name)
			if !BlackFlag {
				plugin_log.Info("丢弃物品: %s,物品Id: %s,物品数量: %d", Article.Name, Article.ID, Article.ArticleCount)
				plugin_sdk.ArticleServiceInstance.DropArticle(Article.ID, Article.ArticleCount)
			}
		}
	}
}

func Convert2DropArticleTimerConfig(OriginConfig interface{}) (*model.DropArticleTimerConfig, error) {
	Config := &model.DropArticleTimerConfig{}
	input, err := json.Marshal(OriginConfig)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(input, Config)
	if err != nil {
		return nil, err
	}
	return Config, nil
}
