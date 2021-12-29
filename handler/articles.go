package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"practice/blog/article/storage"
	"strconv"

	apb "practice/blog/gunk/v1/article"

	"github.com/gorilla/mux"
)

type ShowArticleData struct {
	Articles       []storage.Articles
	LoggedUsername string
}

type ShowAllArticleData struct {
	Articles       []*storage.Articles
	LoggedUsername string
}

//data to show one article by id
type ShowArticleByIdData struct {
	Article        storage.Articles
	LoggedUsername string
	CheckAuthor    bool //true if the username and createdBy matches
}

type CreateArticleData struct {
	ArticleAuth map[string]error
	Article     storage.Articles
}

type UpdateArticleData struct {
	ArticleAuth    map[string]error
	Article        storage.Articles
	CheckAuthor    bool
	LoggedUsername string
}

//**************************************** show article handler **********************************************

func (s *Server) showArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("show an article")
	data := ShowAllArticleData{}
	uname, ok := s.CheckLoggedIn(r)
	if ok {
		data.LoggedUsername = uname
	}
	res, err := s.asrv.GetAllArticles(context.Background(), &apb.GetAllArticleRequest{})
	if err != nil {
		log.Println("error for getting all articles from database ", err)
	}
	//convert  apb.Articles to storage article
	ppfs := []*storage.Articles{}
	for _, pf := range res.Articles {
		ppfs = append(ppfs, ProtoToStorsge(pf))
	}
	if err != nil {
		log.Println("error getting to core.Get()")
	}
	data.Articles = ppfs

	err = s.templates.ExecuteTemplate(w, "show-article.html", data)
	if err != nil {
		log.Println("error executing show-article template", err)
	}
}

func (s *Server) showArticleByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is full showArticleByID ")

	params := mux.Vars(r)
	articleID := params["id"]

	integerValue, err := strconv.Atoi(articleID)
	if err != nil {
		log.Print("error converting a articleID to integer", err)
	}
	ID := int32(integerValue)

	req := &apb.GetArticleRequest{
		ID: ID,
	}
	res, err := s.asrv.GetArticle(context.Background(), req)
	if err != nil {
		log.Println("error getting the article by ID ", err)
	}

	data := ShowArticleByIdData{
		Article: storage.Articles{
			Title:       res.Article.Title,
			Description: res.Article.Description,
			Author:      res.Article.Author,
			ID:          ID,
		},
		LoggedUsername: "",
		CheckAuthor:    false,
	}
	//********** check the retrieved username and session username to match user and author
	Username, ok := s.CheckLoggedIn(r)
	if ok {
		if Username == res.Article.Author {
			data.CheckAuthor = true //author matched
		}
	}
	data.LoggedUsername = Username

	s.templates.ExecuteTemplate(w, "index-articleT.html", data)
}

//**************************************** create article handler **********************************************

func (s *Server) createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create an article")

	//check if logged in
	_, ok := s.CheckLoggedIn(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		err := s.templates.ExecuteTemplate(w, "create-article.html", nil)
		log.Println("error loading create-articale page ", err)
	}

}

func (s *Server) createArticlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create article post")
	newArticle := storage.Articles{}

	//get the username from session
	sess, err := s.session.Get(r, "user-session")
	if err != nil {
		log.Println("error getting session in create-article: ", err)
	}
	if username := sess.Values["user_name"].(string); len(username) > 0 {
		newArticle.Author = username
		newArticle.UserID = sess.Values["user_ID"].(int32)
	}

	r.ParseForm()
	err = s.decoder.Decode(&newArticle, r.PostForm)
	if err != nil {
		log.Println("database connection problem: ", err)
	}

	req := &apb.CreateArticleRequest{
		Article: StorageToProto(newArticle),
	}
	res, err := s.asrv.CreateArticle(context.Background(), req)
	if err != nil {
		fmt.Println("error creating article: ", err)
	}
	fmt.Printf("response from create article: %#v", res)
	http.Redirect(w, r, "/show-article", http.StatusFound)

}

//**************************************** update article handler **********************************************

func (s *Server) updateArticleGet(w http.ResponseWriter, r *http.Request) {
	var ID int32
	var Uname string
	var ok bool
	data := UpdateArticleData{}

	//get article from id
	params := mux.Vars(r)
	articleID := params["id"]
	integerValue, err := strconv.Atoi(articleID)
	if err != nil {
		log.Print("error converting a articleID to integer", err)
	}

	ID = int32(integerValue)
	log.Printf("ID from params in integer: %#v", ID)

	req := &apb.GetArticleRequest{
		ID: ID,
	}
	res, err := s.asrv.GetArticle(context.Background(), req)
	if err != nil {
		log.Println("error getting the article by ID ", err)
	}

	article := ProtoToStorsge(res.Article)
	data.Article = *article

	if err != nil {
		log.Println("error getting the article by ID", err)
	}
	//check if the user logged in and match author with logged user
	Uname, ok = s.CheckLoggedIn(r)
	if ok && Uname != article.Author {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if Uname == article.Author {
		data.CheckAuthor = true
	}
	data.LoggedUsername = Uname
	s.templates.ExecuteTemplate(w, "update-article.html", data)

}

func (s *Server) updateArticlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("update article post")
	article := storage.Articles{}

	//get article from id
	params := mux.Vars(r)
	articleID := params["id"]
	intValue, err := strconv.Atoi(articleID)
	if err != nil {
		log.Printf("error converting article id from string to integer")
	}
	Int32Value := int32(intValue)
	r.ParseForm()
	err = s.decoder.Decode(&article, r.PostForm)
	if err != nil {
		log.Println("error decoding form into struct at create-article ", err)
	}
	article.ID = Int32Value

	a := StorageToProto(article)
	//update article in database
	req := &apb.UpdateArticleRequest{
		Article: a,
	}
	_, err = s.asrv.UpdateArticle(context.Background(), req)
	if err != nil {
		log.Println("error updating article data", err)
	}
	http.Redirect(w, r, "/show-article/"+articleID, http.StatusFound)

}

//**************************************** delete article handler **********************************************
func (s *Server) deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete the article")
	params := mux.Vars(r)
	articleID := params["id"]
	log.Println("id : ", articleID)
	integerValue, err := strconv.Atoi(articleID)
	if err != nil {
		log.Print("error converting a string to integer", err)
	}

	//convert articleID type to int32from string
	ID := int32(integerValue)
	req := &apb.DeleteArticleRequest{
		ID: ID,
	}

	//delete article from db using id
	_, err = s.asrv.DeleteArticle(context.Background(), req)
	if err != nil {
		log.Println("error deleting the article, ", err)
	}
	log.Println("error deleting row with given id: ", err)
	http.Redirect(w, r, "/show-article", http.StatusFound)

}
