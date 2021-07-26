package postgres

import (
	"Blog/storage"
	"log"
)

const (
	getAllArticles = `SELECT id,title, description, uid,username,created_at FROM articles`

	createAnArticle = `INSERT INTO articles VALUES(DEFAULT, :title, :description,  :uid,:username, now(), now())
	RETURNING id`

	getAnArticle = `SELECT * FROM articles WHERE id=$1`
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

	//stmnt, err:=s.Db.PrepareNamed(getAllArticles)
	err := s.Db.Select(&articles, getAllArticles)

	return articles, err

}

func (s *StoreDB) ShowIndexedArticle(id int32) (storage.Articles, error) {
	var article storage.Articles
	err := s.Db.Get(&article, getAnArticle, id)

	return article, err

}
