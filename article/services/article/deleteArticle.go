package article

import (
	"context"
	"fmt"
	apb "practice/blog/gunk/v1/article"
)

func (s *Svc) DeleteArticle(ctx context.Context, req *apb.DeleteArticleRequest) (*apb.DeleteArticleResponse, error) {
	//validation if needed

	err := s.core.Delete(context.Background(), req.GetID())
	if err != nil {
		fmt.Println("error getting to core.Update()")
	}
	return &apb.DeleteArticleResponse{}, nil
}
