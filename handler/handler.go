package handler

import (
	"Blog/storage/postgres"
	"log"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}

func CheckEmptyString(str string) bool {
	var fl bool
	if len(str) != 0 {
		fl = true
	} else {
		fl = false
	}
	return fl
}

type (
	Server struct {
		templates *template.Template
		store     *postgres.StoreDB
		decoder   *schema.Decoder
		session   *sessions.CookieStore
	}
)

func NewServer(store *postgres.StoreDB, decoder *schema.Decoder, session *sessions.CookieStore) (*mux.Router, error) {
	r := mux.NewRouter()

	s := &Server{

		//templates: template.Must(template.ParseGlob("assets/*/*.html")),
		store:   store,
		decoder: decoder,
		session: session,
	}

	//parse template with help of a function
	err := s.parseTemplates()
	CheckError("template parsing error: ", err)

	//define all routes here
	r.HandleFunc("/", s.indexHandler).Methods("GET")
	r.HandleFunc("/signup", s.signupGetHandler).Methods("GET")
	r.HandleFunc("/signupPost", s.signupPostHandler).Methods("POST")
	r.HandleFunc("/login", s.loginGetHandler).Methods("GET")
	r.HandleFunc("/loginPost", s.loginPostHandler).Methods("POST")
	r.HandleFunc("/create-article", s.createArticle).Methods("GET")
	r.HandleFunc("/show-article", s.showArticle).Methods("GET")
	r.HandleFunc("/logout", s.logOut).Methods("GET")
	return r, nil

}

func (s *Server) parseTemplates() error {
	templates := template.New("temp").Funcs(template.FuncMap(sprig.FuncMap()))
	tmpls, err := templates.ParseGlob("assets/*/*.html")
	if err != nil {
		return err
	}

	s.templates = tmpls
	return nil

}

/*
func reverse(str string) string {
	n := len(str)
	runes := make([]rune, n)
	for _, rune := range str {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
*/