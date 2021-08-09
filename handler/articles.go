package handler

import (
	"Blog/storage"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ShowArticleData struct {
	Articles       []storage.Articles
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
	data := ShowArticleData{}
	uname, ok := s.CheckLoggedIn(r)
	if ok {
		data.LoggedUsername = uname
	}

	AllArticles, err := s.store.ShowAllArticles()
	CheckError("error for getting all articles from database ", err)

	data.Articles = AllArticles

	err = s.templates.ExecuteTemplate(w, "show-article.html", data)
	CheckError("error executing show-article template", err)
}

func (s *Server) showArticleByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is full showArticleByID ")
	var data ShowArticleByIdData

	////******get article id from url params and get the article from db
	params := mux.Vars(r)
	articleID := params["id"]

	//convert articleID type to int32from string
	ID := int32(ConvertStringtoInt(articleID))

	//save article to a struct for sending to template
	articleByID, err := s.store.GetIndexedArticle(ID)
	CheckError("error getting the article by ID ", err)

	//********** check the retrieved username and session username to match user and author
	Username, ok := s.CheckLoggedIn(r)
	if ok {
		if Username == articleByID.Author {
			data.CheckAuthor = true //author matched
		}
	}
	data.Article = articleByID
	data.LoggedUsername = Username

	s.templates.ExecuteTemplate(w, "index-articleT.html", data)
}

//**************************************** create article handler **********************************************

func (s *Server) createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create an article")
	//data := ArticleTempData{}

	//check if logged in
	_, ok := s.CheckLoggedIn(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		err := s.templates.ExecuteTemplate(w, "create-article.html", nil)
		CheckError("error loading create-articale page ", err)
	}

}

func (s *Server) createArticlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create article post")
	//data := ArticleTempData{}
	newArticle := storage.Articles{}

	//get the username from session
	sess, err := s.session.Get(r, "user-session")
	CheckError("error getting session in create-article: ", err)
	if username := sess.Values["user_name"].(string); len(username) > 0 {
		newArticle.Author = username
		newArticle.UserID = sess.Values["user_ID"].(int32)
	}

	//get values from input fields (title and description)
	r.ParseForm()
	err = s.decoder.Decode(&newArticle, r.PostForm)
	CheckError("error decoding form into struct at create-article ", err)

	fmt.Printf("articles from html form: %v", newArticle)

	//******************store article in database***************************************
	_, DBerr := s.store.CreateArticle(newArticle)
	CheckError("error inserting article data", DBerr)

	//data = ArticleTempData{}
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
	//convert articleID type to int32from string
	ID = int32(ConvertStringtoInt(articleID))

	//store returned data from db articles
	articleByID, err := s.store.GetIndexedArticle(ID)
	CheckError("error getting the article by ID", err)

	//check if the user logged in and match author with logged user
	Uname, ok = s.CheckLoggedIn(r)
	if !ok || Uname != articleByID.Author {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if Uname == articleByID.Author {
		data.CheckAuthor = true

	}
	//check if the user and author is same

	data.Article = articleByID
	data.LoggedUsername = Uname
	s.templates.ExecuteTemplate(w, "update-article.html", data)

}

func (s *Server) updateArticlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("update article post")
	data := UpdateArticleData{}
	article := storage.Articles{}

	//****** check if logged in and user matches the author
	//get article from id
	params := mux.Vars(r)
	articleID := params["id"]

	//convert articleID type to int32from string
	ID := int32(ConvertStringtoInt(articleID))
	article.ID = ID

	//store returned data from db articles
	articleByID, err := s.store.GetIndexedArticle(ID)
	CheckError("error getting the article by ID", err)

	username, ok := s.CheckLoggedIn(r)

	if !ok || username != articleByID.Author {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else if username == articleByID.Author {
		data.CheckAuthor = true
	}

	//get values from input fields (title and description)
	r.ParseForm()
	err = s.decoder.Decode(&article, r.PostForm)
	CheckError("error decoding form into struct at create-article ", err)

	//******************store article in database***************************************
	DBerr := s.store.UpdateIndexedArticle(article)
	CheckError("error updating article data", DBerr)

	redirectedUrl := "/show-article/" + string(articleID)
	http.Redirect(w, r, redirectedUrl, http.StatusFound)

}

//**************************************** delete article handler **********************************************
func (s *Server) deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete the article")
	var article storage.Articles
	//****** check if logged in and user matches the author
	//get article from id
	params := mux.Vars(r)
	articleID := params["id"]

	//convert articleID type to int32from string
	ID := int32(ConvertStringtoInt(articleID))
	article.ID = ID

	//delete article from db using id
	err := s.store.DeleteArticleByID(ID)
	CheckError("error deleting row with given id: ", err)

	http.Redirect(w, r, "/show-article", http.StatusFound)

}
