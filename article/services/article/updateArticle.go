package article

import (
	"context"
	"fmt"
	apb "practice/blog/gunk/v1/article"
)

func (s *Svc) UpdateArticle(ctx context.Context, req *apb.UpdateArticleRequest) (*apb.UpdateArticleResponse, error) {
	//validation if needed
	ppfs := ProtoToStorage(req.Article)
	fmt.Printf("printing req after protoToStorage(): %#v", ppfs)

	err := s.core.Update(context.Background(), ppfs)
	if err != nil {
		fmt.Println("error getting to core.Update()")
	}
	return &apb.UpdateArticleResponse{}, nil
}
