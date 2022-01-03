package handler

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"practice/blog/article/storage"
	"practice/blog/article/storage/postgres"
	"text/template"
	"time"

	apb "practice/blog/gunk/v1/article"

	"github.com/Masterminds/sprig/v3"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

type (
	Server struct {
		templates *template.Template
		store     *postgres.StoreDB
		decoder   *schema.Decoder
		session   *sessions.CookieStore
		asrv      apb.ArticleClient
	}
)

func NewServer(store *postgres.StoreDB, decoder *schema.Decoder, session *sessions.CookieStore, ab apb.ArticleClient) (*mux.Router, error) {
	r := mux.NewRouter()

	s := &Server{
		store:   store,
		decoder: decoder,
		session: session,
		asrv:    ab,
	}

	//parse template with help of a function
	err := s.parseTemplates()
	if err != nil {
		log.Println("template parsing error: ", err)
	}

	//define all routes here
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../assets/"))))

	r.HandleFunc("/", s.indexHandler).Methods("GET")
	r.HandleFunc("/signup", s.signupGetHandler).Methods("GET")
	r.HandleFunc("/signupPost", s.signupPostHandler).Methods("POST")
	r.HandleFunc("/login", s.loginGetHandler).Methods("GET")
	r.HandleFunc("/loginPost", s.loginPostHandler).Methods("POST")
	r.HandleFunc("/create-article", s.createArticle).Methods("GET")
	r.HandleFunc("/create-article", s.createArticlePost).Methods("POST")
	r.HandleFunc("/show-article", s.showArticle).Methods("GET")
	r.HandleFunc("/show-article/{id}", s.showArticleByID).Methods("GET")
	r.HandleFunc("/show-article/{id}", s.showArticleByID).Methods("POST")

	r.HandleFunc("/update-article/{id}", s.updateArticleGet).Methods("GET")
	r.HandleFunc("/update-article/{id}", s.updateArticlePost).Methods("POST")
	r.HandleFunc("/delete-article/{id}", s.deleteArticle).Methods("POST")

	r.HandleFunc("/logout", s.logOut).Methods("GET")
	return r, nil

}

func (s *Server) parseTemplates() error {
	templates := template.New("temp").Funcs(template.FuncMap{
		"formatDateNull": func(i interface{}, format, nullDefault string) (string, error) {
			var t time.Time
			switch v := i.(type) {
			case time.Time:
				t = v
			case *time.Time:
				if v == nil {
					return nullDefault, nil
				}
				t = *v
			case sql.NullString:
				if !v.Valid {
					return nullDefault, nil
				}
				// tt, err := tryParseDate(v.String)
				// if err != nil {
				// 	return "", err
				// }
				// t = tt
			case string:
				tt, err := time.Parse("2006-01-02T15:04:05Z", i.(string))
				if err != nil {
					return "", err
				}
				t = tt
			default:
				return "", errors.New("unknown type of date")
			}
			formatted := t.Format(format)
			return formatted, nil
		},
	}).Funcs(template.FuncMap(sprig.FuncMap()))
	tmpls, err := templates.ParseGlob("../assets/*/*.html")
	if err != nil {
		return err
	}

	s.templates = tmpls
	return nil

}

func ProtoToStorsge(spf *apb.Articles) *storage.Articles {
	ppf := &storage.Articles{
		ID:          spf.ID,
		Title:       spf.Title,
		Description: spf.Description,
		Author:      spf.Author,
		UserID:      spf.UserID,
		CreatedAt:   spf.CreatedAt.AsTime(),
		UpdatedAt:   spf.CreatedAt.AsTime(),
	}
	return ppf
}

func StorageToProto(spf storage.Articles) *apb.Articles {
	ppf := &apb.Articles{
		ID:          spf.ID,
		Title:       spf.Title,
		Description: spf.Description,
		Author:      spf.Author,
		UserID:      spf.UserID,
		CreatedAt:   tspb.New(spf.CreatedAt),
		UpdatedAt:   tspb.New(spf.UpdatedAt),
	}
	return ppf
}
