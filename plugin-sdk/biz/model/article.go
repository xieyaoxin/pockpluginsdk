package model

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
