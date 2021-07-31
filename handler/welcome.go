package handler

import (
	"net/http"
)

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	username, loggedIn := s.CheckLoggedIn(r)
	User := struct {
		Uname    string
		LoggedIn bool
	}{
		Uname:    username,
		LoggedIn: loggedIn,
	}
	s.templates.ExecuteTemplate(w, "index.html", User)
}
