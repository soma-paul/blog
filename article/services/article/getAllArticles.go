package article

import (
	"context"
	"log"
	apb "practice/blog/gunk/v1/article"
)

func (s *Svc) GetAllArticles(ctx context.Context, req *apb.GetAllArticleRequest) (*apb.GetAllArticleResponse, error) {
	//validation if needed

	articles, err := s.core.GetAll(ctx)
	ppfs := []*apb.Articles{}
	for _, pf := range articles {
		ppfs = append(ppfs, StorageToProto(pf))
	}
	if err != nil {
		log.Println("error getting to core.Get()")
	}
	return &apb.GetAllArticleResponse{
		Articles: ppfs,
	}, nil
}
