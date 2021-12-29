package article

import (
	"context"
	"log"
	apb "practice/blog/gunk/v1/article"
)

func (s *Svc) GetArticle(ctx context.Context, req *apb.GetArticleRequest) (*apb.GetArticleResponse, error) {
	//validation if needed
	id := int32(req.GetID())
	article, err := s.core.Get(ctx, id)
	a := StorageToProto(article)
	if err != nil {
		log.Println("error getting to core.Get()")
	}
	return &apb.GetArticleResponse{
		Article: a,
	}, nil
}
