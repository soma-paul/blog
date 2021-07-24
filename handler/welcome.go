package handler

import (
	"net/http"
)

//these will be used to define corresponding handlerFunction
//template parsing will be done here too.
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	// db := dbConn()
	// defer db.Close()
	// fmt.Println("hello!")
	s.templates.ExecuteTemplate(w, "index.html", nil)
}
