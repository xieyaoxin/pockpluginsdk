package kdhs

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/log"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/model"
	util "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/utils"
	"strconv"
	"strings"
)

var ArticleRepositoryImpl4KDHSInstance = &articleRepositoryImpl4KDHS{}

type (
	articleRepositoryImpl4KDHS struct {
	}
)

func (*articleRepositoryImpl4KDHS) UseArticle(articleId string) error {
	params := util.InitParam()
	params["id"] = articleId
	result := CallServerGetInterface("function/usedProps.php", params)
	log.Info("使用物品 id : %s 操作结果: %s", articleId, result)
	return nil
}
func (rep *articleRepositoryImpl4KDHS) GetArticles(name string) ([]*model.Article, error) {
	parmas := util.InitParam()
	parmas["style"] = "1"
	result := CallServerGetInterface("function/getBag.php", parmas)

	lines := strings.Split(result, "\n")
	articleList := []*model.Article{}
	for lineNumber := range lines {
		line := lines[lineNumber]
		if strings.Contains(line, name) && strings.Contains(line, "bid") {
			id := strings.Replace(strings.Split(strings.Split(line, "bid")[1], ";")[0], "=", "", 1)
			//article := rep.GetArticleDetail(id)
			article := model.Article{
				ID: id,
			}
			articleName := strings.Split(strings.Replace(line, "<", ">", -1), ">")[2]
			article.Name = articleName

			articleList = append(articleList, &article)

			articleTypeLine := lines[lineNumber+1]
			articleType := strings.Split(strings.Replace(articleTypeLine, "<", ">", -1), ">")[2]
			article.ArticleType = articleType

			articleCountLine := lines[lineNumber+2]
			articleCount, _ := strconv.ParseInt(strings.Split(strings.Replace(articleCountLine, "<", ">", -1), ">")[2], 10, 64)
			article.ArticleCount = int(articleCount)
		}

	}
	//log.Info("查询到物品列表：%s %s", name, util.ListToJson(articleList))

	return articleList, nil
}

func (*articleRepositoryImpl4KDHS) GetArticleDetail(articleId string) model.Article {
	params := util.InitParam()
	//id=414880&bid=0&sign=1&type=2
	params["id"] = articleId
	params["bid"] = "0"
	params["sign"] = "1"
	params["type"] = "2"
	result := CallServerGetInterface("function/getPropsInfo.php", params)
	article := model.Article{ID: articleId}
	article.Sellable = !strings.Contains(result, "不可交易")
	return article
}
