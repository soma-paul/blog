package handler

import (
	"Blog/storage"
	"fmt"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type LoginTempData struct {
	User         storage.Users
	UserAuth     map[string]string //registered email, not empty field, correct password check
	LoginSuccess bool
	LoggedIn     bool
}

func ValidateLogin(user storage.Users) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required.Error("Please give your email address.")),
		validation.Field(&user.Password, validation.Required.Error("please give your password here.")),
	)
}

func (s *Server) loginGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
	data := LoginTempData{}

	//check if already logged in
	newSession, err := s.session.Get(r, "user-session")
	CheckError("error getting session from login GET ", err)
	if len(newSession.Values) > 0 {
		data.UserAuth = map[string]string{}
		data.UserAuth["Logout"] = "you are already logged in. "
		data.LoggedIn = true
	}

	err = s.loadLogin(w, r, data)
	CheckError("error loding get log in form ", err)

}

func (s *Server) loginPostHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("login post")
	var usrForm storage.Users
	var data LoginTempData
	r.ParseForm()
	err := s.decoder.Decode(&usrForm, r.PostForm)
	CheckError("error decoding log in form: ", err)

	varErr := map[string]string{}
	validationErr := ValidateLogin(usrForm)
	if validationErr != nil {
		if e, ok := validationErr.(validation.Errors); ok {
			if len(e) > 0 {

				for key, value := range e {
					varErr[key] = value.Error()
				}
			}
			//data.UserAuth = varErr
		}

	}
	var passwordErr error
	userDB := s.store.UserAuth(usrForm.Email) //retrieves id, username, password for given email

	if userDB.ID == 0 {
		varErr["EmailNotValid"] = "this email is not valid"

	} else {
		//check password
		if usrForm.Password != "" {
			passwordErr = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(usrForm.Password))
			if passwordErr != nil {
				log.Printf("error for matching password: %v", passwordErr)
				varErr["PassNotMatched"] = "Invalid password, try again!"
			}
		}

	}
	//fmt.Printf("password err: %v, validation error: %v", passwordErr, varErr)
	if len(varErr) == 0 {
		//create session
		newSession, er := s.session.Get(r, "user-session")
		fmt.Printf("priniting session: %T, %+v", newSession, newSession)
		CheckError("error creating session", er)

		//set values like username and userid depending on the user
		newSession.Values["user_name"] = userDB.Username
		newSession.Values["user_ID"] = userDB.ID
		newSession.Values["isAdmin"] = userDB.IsAdmin
		err = newSession.Save(r, w)
		CheckError("error saving session: ", err)
		data.LoginSuccess = true
		usrForm = storage.Users{}
	}

	data.User = usrForm
	data.UserAuth = varErr

	err = s.loadLogin(w, r, data)
	CheckError("error loading post log in form ", err)

}

func (s *Server) loadLogin(w http.ResponseWriter, r *http.Request, data LoginTempData) error {
	err := s.templates.ExecuteTemplate(w, "login.html", data)
	return err
}

func (s *Server) logOut(w http.ResponseWriter, r *http.Request) {
	session, err := s.session.Get(r, "user-session")
	CheckError("error in getting session in logout ", err)
	delete(session.Values, "user_ID")
	delete(session.Values, "user_name")
	delete(session.Values, "isAdmin")
	err = session.Save(r, w)
	CheckError("error saving session after log out ", err)

	http.Redirect(w, r, "/login", 307)
}
