package model

type Article struct {
	*PockBaseModel
	ID           string
	Name         string
	ArticleType  string
	Sellable     bool
	ArticleCount int
}
