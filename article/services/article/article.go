package article

import (
	"context"
	"practice/blog/article/storage"
	apb "practice/blog/gunk/v1/article"

	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

type articleCoreStore interface {
	Get(ctx context.Context, id int32) (*storage.Articles, error)
	GetAll(ctx context.Context) ([]*storage.Articles, error)
	CreateArticle(ctx context.Context, article storage.Articles) (int32, error)
	Update(ctx context.Context, article storage.Articles) error
	Delete(ctx context.Context, id int32) error
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

func StorageToProto(spf *storage.Articles) *apb.Articles {
	ppf := &apb.Articles{
		ID:          spf.ID,
		Title:       spf.Title,
		Description: spf.Description,
		Author:      spf.Author,
		UserID:      spf.UserID,
		CreatedAt:   tspb.New(spf.CreatedAt),
		UpdatedAt:   tspb.New(spf.UpdatedAt),
	}
	return ppf
}
