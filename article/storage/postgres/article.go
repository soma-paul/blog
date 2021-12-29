package postgres

import (
	"fmt"
	"log"
	"practice/blog/article/storage"
)

const (
	getAllArticles = `SELECT id,title, description, uid,username, created_at FROM articles`
)

func (s *StoreDB) CreateArticle(data storage.Articles) (int32, error) {
	const createAnArticle = `INSERT INTO articles VALUES(DEFAULT, :title, :description,  :uid, :username, now(), now())
	RETURNING id`
	stmnt, err := s.Db.PrepareNamed(createAnArticle)
	if err != nil {
		log.Fatalf("error making named statement for creating article %v", err)
	}
	var id int32
	err = stmnt.Get(&id, data)
	return id, err

}

func (s *StoreDB) ShowAllArticles() ([]*storage.Articles, error) {
	var articles []*storage.Articles
	err := s.Db.Select(&articles, getAllArticles)
	fmt.Printf("printing all articles: %#v", articles[0].CreatedAt)

	return articles, err

}

func (s *StoreDB) GetIndexedArticle(id int32) (*storage.Articles, error) {
	const getAnArticle = `SELECT id,title, description, uid,username,created_at FROM articles WHERE id=$1`
	var article storage.Articles
	err := s.Db.Get(&article, getAnArticle, id)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (s *StoreDB) UpdateIndexedArticle(data storage.Articles) error {
	const UpdateAnArticle = `UPDATE articles SET title=:title, description=:description, updated_at=now() WHERE id=:id`
	result, err := s.Db.NamedExec(UpdateAnArticle, data)
	fmt.Printf("prinint db updating result: %T %+v", result, result)
	return err
}

func (s *StoreDB) DeleteArticleByID(id int32) error {
	const DeleteAnArticle = `DELETE FROM articles WHERE id=$1`
	_, err := s.Db.Query(DeleteAnArticle, id)
	return err
}
