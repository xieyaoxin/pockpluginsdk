package repository

type ArticleRepository interface {
	UseArticle(articleId string) error
}
