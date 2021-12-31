package article

import (
	"context"
	"fmt"
	"practice/blog/article/storage"
	apb "practice/blog/gunk/v1/article"
	"time"
)

func (s *Svc) CreateArticle(ctx context.Context, req *apb.CreateArticleRequest) (*apb.CreateArticleResponse, error) {
	//validation if needed

	ppfs := ProtoToStorage(req.Article)
	id, err := s.core.CreateArticle(context.Background(), ppfs)
	if err != nil {
		fmt.Println("error getting to core.CreateArticle()")
	}
	return &apb.CreateArticleResponse{
		ID: id,
	}, nil
}

func ProtoToStorage(spf *apb.Articles) storage.Articles {
	ppf := storage.Articles{
		ID:          spf.ID,
		Title:       spf.Title,
		Description: spf.Description,
		Author:      spf.Author,
		UserID:      spf.UserID,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	return ppf
}
