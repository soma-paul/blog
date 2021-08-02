package handler

import (
	"Blog/storage"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

func validateSignup(u storage.SignupUser) error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required.Error("firstname is required!")),
		validation.Field(&u.LastName, validation.Required.Error("lastname is required!")),
		validation.Field(&u.Username, validation.Required.Error("username is required!")),
		validation.Field(&u.Email, validation.Required.Error("email is required!")),
		validation.Field(&u.Password, validation.Required.Error("Password is required!")),
	)
}

type UserTempData struct {
	NewUser      storage.SignupUser
	FormValidate map[string]string
	Success      bool
}

func (s *Server) signupGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup")
	data := UserTempData{}
	s.loadTemplate(w, r, data)
}

func (s *Server) signupPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup post")
	var usr storage.SignupUser
	var data UserTempData

	s.decoder.IgnoreUnknownKeys(true)

	err := r.ParseForm()
	CheckError("form parsing error:", err)

	//decoded form value into usr stuct
	err = s.decoder.Decode(&usr, r.PostForm)
	CheckError("error decoding schema from sign up user: ", err)

	//fmt.Printf("user check before assigning to formdata: %v", usr)

	data.FormValidate = make(map[string]string)

	//check uniqueness of email and username
	eID, uID := s.store.UniqueEmailUname(usr.Email)
	varErr := map[string]string{}

	//validate data
	err = validateSignup(usr)

	//****if validationError== nil and email, username is unique, then execute db query to insert row

	if err != nil || eID != 0 || uID != 0 {

		if e, ok := err.(validation.Errors); ok {
			if len(e) > 0 {
				for key, value := range e {
					varErr[key] = value.Error()
				}

			}

		}
		if eID == 0 {
			varErr["DuplicateEmail"] = "Sorry! Email address already in use"
		}
		if uID == 0 {
			varErr["DuplicateUsername"] = "Sorry! Username is already in use, May be try another one?"
		}
		data.FormValidate = varErr
		data.NewUser = usr

	} else {
		//hash the password
		hash, er := bcrypt.GenerateFromPassword([]byte(usr.Password), 10)
		CheckError("error generating hash ", er)
		hashString := string(hash)
		usr.Password = hashString //save hashed password in string type

		//save user to the
		var id int32
		er, id = s.store.CreateUser(usr) //store the user in users table
		CheckError("error executing query for creating an user: ", err)
		fmt.Println("returned id: ", id)

		data = UserTempData{} //clear the data if post is done successfully.
		data.Success = true
	}

	s.loadTemplate(w, r, data)

	//if validation err==nil && database job is done successfully, then execute this
	http.Redirect(w, r, "/signup", 302)

}

func (s *Server) loadTemplate(w http.ResponseWriter, r *http.Request, data UserTempData) {
	err := s.templates.ExecuteTemplate(w, "signupT.html", data)
	CheckError("error parsing signup.html: ", err)
}
