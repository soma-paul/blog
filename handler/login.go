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
	User     storage.Users
	UserAuth map[string]string //registered email, not empty field, correct password check
	Username string
}

func ValidateLogin(user storage.Users) map[string]string {
	validationErr := make(map[string]string)
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required.Error("Please give your email address.")),
		validation.Field(&user.Password, validation.Required.Error("please give your password here.")),
	)
	if err != nil {
		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					validationErr[key] = value.Error()
				}
			}
		}

	}

	return validationErr
}

func (s *Server) loginGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
	data := LoginTempData{}

	//check if logged in
	u, loggedIn := s.CheckLoggedIn(r)
	if loggedIn {
		data.UserAuth = map[string]string{}
		data.UserAuth["Logout"] = "you are already logged in. "
		data.Username = u
	}

	fmt.Printf("*********************data: %+v", data)

	/*
		newSession, err := s.session.Get(r, "user-session")
		CheckError("error getting session from login GET ", err)
		if len(newSession.Values) > 0 {
			data.UserAuth = map[string]string{}
			data.UserAuth["Logout"] = "you are already logged in. "
			data.LoggedIn = true
		}
	*/

	err := s.loadLogin(w, data)
	CheckError("error loding get log in form ", err)

}

func (s *Server) loginPostHandler(w http.ResponseWriter, r *http.Request) {
	var usrForm storage.Users
	var data LoginTempData
	r.ParseForm()
	err := s.decoder.Decode(&usrForm, r.PostForm)
	CheckError("error decoding log in form: ", err)

	//**********************Validate Login Form********************************************
	data.UserAuth = map[string]string{}    //initializing map
	data.UserAuth = ValidateLogin(usrForm) //validating log in form
	/*
		if validationErr != nil {
			if e, ok := validationErr.(validation.Errors); ok {
				if len(e) > 0 {

					for key, value := range e {
						varErr[key] = value.Error()
					}
				}
			}

		}
	*/
	var passwordErr error
	userDB := s.store.UserAuth(usrForm.Email) //retrieves id, username, password for given email

	//check if entered email is registered
	if userDB.ID == 0 {
		data.UserAuth["EmailNotValid"] = "this email is not valid"

	} else {
		//check password
		if usrForm.Password != "" {
			passwordErr = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(usrForm.Password))
			if passwordErr != nil {
				log.Printf("error for matching password: %v", passwordErr)
				data.UserAuth["PassNotMatched"] = "Invalid password, try again!"
			}
		}

	}
	//fmt.Printf("password err: %v, validation error: %v", passwordErr, varErr)
	if len(data.UserAuth) == 0 {
		//create session
		err = s.CreateSession(w, r, userDB)
		CheckError("error creating and saving session. ", err)

		//redirect to homepage if logged in successfully
		http.Redirect(w, r, "/", 302)

		usrForm = storage.Users{}

	}

	data.User = usrForm

	err = s.loadLogin(w, data)
	CheckError("error loading post log in form ", err)

}

func (s *Server) loadLogin(w http.ResponseWriter, data LoginTempData) error {
	err := s.templates.ExecuteTemplate(w, "login.html", data)
	return err
}

func (s *Server) logOut(w http.ResponseWriter, r *http.Request) {
	err := s.DeleteSession(w, r)
	CheckError("Error for saving session after deleting values ", err)

	http.Redirect(w, r, "/login", 307)
}
