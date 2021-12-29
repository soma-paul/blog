package article

import (
	"practice/blog/article/storage"
)

type articleStore interface {
	GetIndexedArticle(id int32) (*storage.Articles, error)
	ShowAllArticles() ([]*storage.Articles, error)
	CreateArticle(data storage.Articles) (int32, error)
	UpdateIndexedArticle(data storage.Articles) error
	DeleteArticleByID(id int32) error
}

type CoreSvc struct {
	artStr articleStore
}

func NewCoreSvc(astr articleStore) *CoreSvc {
	return &CoreSvc{astr}
}
