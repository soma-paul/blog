package handler

import (
	"net/http"
)

//these will be used to define corresponding handlerFunction
//template parsing will be done here too.
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
