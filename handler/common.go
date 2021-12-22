package handler

import (
	"fmt"
	"log"
	"net/http"
	"practice/blog/article/storage"
	"strconv"

	"github.com/gorilla/sessions"
)

func ConvertStringtoInt(s string) int {
	integerValue, err := strconv.Atoi(s)
	CheckError("error converting a string to integer", err)
	return integerValue
}

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}

func (s *Server) CheckLoggedIn(r *http.Request) (string, bool) {
	var Uname string
	var LoggedIn bool
	session := s.GetSession(r)
	if len(session.Values) > 0 {

		if username := session.Values["user_name"]; username != nil {
			Uname = username.(string)
			LoggedIn = session.Values["logged_in"].(bool)
		}
	}
	fmt.Printf("username common: %v, loggedIn: %v", Uname, LoggedIn)
	return Uname, LoggedIn
}

func (s *Server) GetSession(r *http.Request) *sessions.Session {
	session, err := s.session.Get(r, "user-session")
	CheckError("error getting session: ", err)
	return session
}

func (s *Server) CreateSession(w http.ResponseWriter, r *http.Request, UserInfo storage.Users) error {
	newSession := s.GetSession(r)

	//set values like username and userid depending on the user
	newSession.Values["user_name"] = UserInfo.Username
	newSession.Values["user_ID"] = UserInfo.ID
	newSession.Values["isAdmin"] = UserInfo.IsAdmin
	newSession.Values["logged_in"] = true
	err := newSession.Save(r, w)
	return err

}

func (s *Server) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session := s.GetSession(r)
	delete(session.Values, "user_ID")
	delete(session.Values, "user_name")
	delete(session.Values, "isAdmin")
	delete(session.Values, "logged_in")
	err := session.Save(r, w)
	return err
}
