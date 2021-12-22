package article

import (
	"context"
	"practice/blog/article/storage"
)

func (cs *CoreSvc) get(ctx context.Context, id string) (storage.Articles, error) {
	return storage.Articles{}, nil
}
