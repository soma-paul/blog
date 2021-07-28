package handler

import (
	"Blog/storage"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ArticleTempData struct {
	Articles []storage.Articles
	Comments map[string]string
}

func (s *Server) showArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("show an article")
	data := ArticleTempData{}
	AllArticles, err := s.store.ShowAllArticles()
	CheckError("error for getting all articles frm database ", err)

	data.Articles = AllArticles
	err = s.templates.ExecuteTemplate(w, "show-article.html", AllArticles)
	CheckError("error executing show-article template", err)
}

func (s *Server) showArticleByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is full view of article")

	params := mux.Vars(r)
	articleID := params["id"]

	//convert articleID type to int32from string
	ID := int32(ConvertStringtoInt(articleID))
	fmt.Println(ID)

	//save article to a struct for sending to template
	article := make([]storage.Articles, 1)
	articleByID, err := s.store.ShowIndexedArticle(ID)
	article[0] = articleByID
	data := ArticleTempData{
		Articles: article,
	}

	fmt.Printf("the article in slice %v", article)
	CheckError("error getting the article by ID ", err)
	s.templates.ExecuteTemplate(w, "index-article.html", data)
}

func (s *Server) showArticleByIDPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := ArticleTempData{}
	data.Comments = make(map[string]string)
	//var AllComments []string
	var newComment string
	newComment = r.FormValue("write-comment")
	username, _ := s.CheckLoggedIn(r)
	data.Comments[username] = newComment
	s.templates.ExecuteTemplate(w, "index-article.html", data)

}

//*****************************For create article **************************************

func (s *Server) createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create an article")
	data := storage.Articles{}

	//check if logged in
	//sess, err := s.session.Get(r, "user-session")
	//CheckError("error getting session in create-article: ", err)
	_, ok := s.CheckLoggedIn(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		s.loadCreateArticle(w, data)
	}

}

func (s *Server) createArticlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("create article post")
	data := storage.Articles{}

	//get the username from session
	sess, err := s.session.Get(r, "user-session")
	CheckError("error getting session in create-article: ", err)
	if username := sess.Values["user_name"].(string); len(username) > 0 {
		data.Author = username
		data.UserID = sess.Values["user_ID"].(int32)
	}

	//get values from input fields (title and description)
	r.ParseForm()
	err = s.decoder.Decode(&data, r.PostForm)
	CheckError("error decoding form into struct at create-article ", err)

	fmt.Printf("articles from html form: %v", data)

	//******************store article in database***************************************
	articleID, DBerr := s.store.CreateArticle(data)
	CheckError("error inserting article data", DBerr)
	fmt.Printf("article is saved to databse with article id: %v", articleID)

	data = storage.Articles{}
	s.loadCreateArticle(w, data)

}

func (s *Server) loadCreateArticle(w http.ResponseWriter, data storage.Articles) {
	err := s.templates.ExecuteTemplate(w, "create-article.html", data)
	CheckError("error loading create-articale page ", err)

}
