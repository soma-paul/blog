package handler

import (
	"fmt"
	"net/http"
)

func (s *Server) showArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("show an article")
	err := s.templates.ExecuteTemplate(w, "show-article.html", nil)
	CheckError("", err)
}

func (s *Server) createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create an article")

	sess, err := s.session.Get(r, "user-session")

	_, ok := sess.Values["user_ID"]

	fmt.Printf("session value: %T, %+v", sess.Values, sess.Values)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	err = s.templates.ExecuteTemplate(w, "create-article.html", nil)
	CheckError("", err)
}
