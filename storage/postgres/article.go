package postgres

import (
	"Blog/storage"
	"fmt"
	"log"
)

const (
	getAllArticles = `SELECT id,title, description, uid,username,created_at FROM articles`

	createAnArticle = `INSERT INTO articles VALUES(DEFAULT, :title, :description,  :uid,:username, now(), now())
	RETURNING id`

	getAnArticle = `SELECT id,title, description, uid,username,created_at FROM articles WHERE id=$1`

	UpdateAnArticle = `UPDATE articles SET title=:title, description=:description WHERE id=:id`
	DeleteAnArticle = `DELETE FROM articles WHERE id=$1`
)

func (s *StoreDB) CreateArticle(data storage.Articles) (int32, error) {
	stmnt, err := s.Db.PrepareNamed(createAnArticle)
	if err != nil {
		log.Fatalf("error making named statement for creating article %v", err)
	}
	var id int32
	err = stmnt.Get(&id, data)
	return id, err

}

func (s *StoreDB) ShowAllArticles() ([]storage.Articles, error) {
	var articles []storage.Articles
	err := s.Db.Select(&articles, getAllArticles)

	return articles, err

}

func (s *StoreDB) GetIndexedArticle(id int32) (storage.Articles, error) {
	var article storage.Articles
	err := s.Db.Get(&article, getAnArticle, id)

	return article, err

}

func (s *StoreDB) UpdateIndexedArticle(data storage.Articles) error {
	result, err := s.Db.NamedExec(UpdateAnArticle, data)
	fmt.Printf("prinint db updating result: %T %+v", result, result)
	return err
}

func (s *StoreDB) DeleteArticleByID(id int32) error {
	_, err := s.Db.Query(DeleteAnArticle, id)
	return err
}
