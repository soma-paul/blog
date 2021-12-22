package article

import (
	"context"
	"fmt"
	"practice/blog/article/storage"
)

type articleStore interface {
	GetIndexedArticle(id int32) (storage.Articles, error)
}

type CoreSvc struct {
	artStr articleStore
}

func NewCoreSvc(astr articleStore) *CoreSvc {
	return &CoreSvc{astr}
}
func (svc *CoreSvc) Get(ctx context.Context, id int32) (storage.Articles, error) {
	fmt.Println("dkrjgkd", id)
	article, err := svc.artStr.GetIndexedArticle(id)
	if err != nil {

	}
	return article, err
}
