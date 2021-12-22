package article

import (
	"context"
	"practice/blog/article/storage"
	apb "practice/blog/gunk/v1/article"
)

type articleCoreStore interface {
	Get(ctx context.Context, id int32) (storage.Articles, error)
}

type Svc struct {
	apb.UnimplementedArticleServer
	core articleCoreStore
}

func NewArticleServer(c articleCoreStore) *Svc {
	return &Svc{
		core: c,
	}
}
