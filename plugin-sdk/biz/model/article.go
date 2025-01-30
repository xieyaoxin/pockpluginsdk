package model

import "fmt"

type Article struct {
	*PockBaseModel
	ID           string
	Name         string
	ArticleType  string
	Sellable     bool
	ArticleCount int
	Pid          string
}

func ArticleSliceRemoveItem(List []*Article, Item *Article) []*Article {
	TempList := []*Article{}
	for _, Item1 := range List {
		if Item != Item1 {
			TempList = append(TempList, Item1)
		}
	}
	return TempList
}

func (ar *Article) GetDetail() string {
	return fmt.Sprintf("物品名称: %s,当前数量: %d,物品ID: %s, 物品类型: %s", ar.Name, ar.ArticleCount, ar.ID, ar.ArticleType)
}
