package article

import (
	"context"
	"fmt"
	"log"
	"practice/blog/article/storage"
)

func (svc *CoreSvc) Get(ctx context.Context, id int32) (*storage.Articles, error) {
	log.Println("article.core.Get()")
	article, err := svc.artStr.GetIndexedArticle(id)
	if err != nil {
		log.Println("error getting article by id, srticle.core.Get()")
	}
	return article, err
}

func (svc *CoreSvc) GetAll(ctx context.Context) ([]*storage.Articles, error) {
	fmt.Println("article.core.GetAll()")
	Articles, err := svc.artStr.ShowAllArticles()
	if err != nil {
		log.Fatal("error creating article: ", err)
		return nil, err
	}
	return Articles, nil

}

func (svc *CoreSvc) CreateArticle(ctx context.Context, article storage.Articles) (int32, error) {
	fmt.Println("article.core.Create()")
	id, err := svc.artStr.CreateArticle(article)
	if err != nil {
		log.Fatal("error creating article: ", err)
	}
	return id, nil
}
func (svc *CoreSvc) Update(ctx context.Context, article storage.Articles) error {
	fmt.Println("article.core.GetAll()")
	_, err := svc.artStr.UpdateIndexedArticle(article)
	if err != nil {
		log.Fatal("error updating article: ", err)

	}
	return nil
}

func (svc *CoreSvc) Delete(ctx context.Context, id int32) error {
	fmt.Println("article.core.GetAll()")
	err := svc.artStr.DeleteArticleByID(id)
	if err != nil {
		log.Println("error deleting the article,  ", err)
	}
	return nil
}
