package postgres

import (
	"fmt"
	"log"
	"practice/blog/article/storage"
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
	const getAllArticles = `SELECT id,title, description, uid,username, created_at FROM articles`
	var articles []*storage.Articles
	err := s.Db.Select(&articles, getAllArticles)
	if err != nil {
		log.Fatalf("error getting all articles, %v", err)
		return nil, err
	}

	return articles, nil

}

func (s *StoreDB) GetIndexedArticle(id int32) (*storage.Articles, error) {
	const getAnArticle = `SELECT id,title, description, uid,username,created_at FROM articles WHERE id=$1`
	var article storage.Articles
	err := s.Db.Get(&article, getAnArticle, id)
	if err != nil {
		log.Printf("error getting an article from list: %v ", err)
		return nil, err
	}
	return &article, nil
}

func (s *StoreDB) UpdateIndexedArticle(data storage.Articles) (storage.Articles, error) {
	const UpdateAnArticle = `UPDATE articles SET title=:title, description=:description, updated_at=now() WHERE id=:id  Returning *`
	var article storage.Articles
	stmnt, err := s.Db.PrepareNamed(UpdateAnArticle)
	if err != nil {
		log.Fatalf("error making named statement for creating article %v", err)
	}
	err = stmnt.Get(&article, data)
	fmt.Printf("prinint db updating result: %T %+v", article, article)
	if err != nil {
		log.Printf("error updating article, %v", err)
		return storage.Articles{}, err

	}
	return article, nil
}

func (s *StoreDB) DeleteArticleByID(id int32) (storage.Articles, error) {
	const DeleteAnArticle = `DELETE FROM articles WHERE id=$1 RETURNING *`
	var article storage.Articles
	err := s.Db.Get(&article, DeleteAnArticle, id)
	if err != nil {
		log.Printf("error deleting article with ID %v error=%v", id, err)
		return storage.Articles{}, err
	}
	return article, nil
}
