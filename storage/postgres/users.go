package postgres

import (
	"Blog/storage"
	"fmt"
	"log"
)

const (
	createUserQuery = `INSERT INTO users
	VALUES(DEFAULT, :first_name, :last_name, :username, :email, :password, true, true, now(), now()) 
	RETURNING id`

	getUser       = `SELECT id,username,password, is_admin FROM users WHERE email=$1`
	getIdForEmail = `SELECT id FROM users WHERE email=$1`
	getIdForUname = `SELECT id FROM users WHERE username=$1`
)

func (s *StoreDB) CreateUser(usr storage.SignupUser) (error, int32) {
	stmt, err := s.Db.PrepareNamed(createUserQuery)
	if err != nil {
		log.Fatalf("error for creating statement %v", err)

	}
	var id int32
	fmt.Printf("passed user to database: %T %+v", usr, usr)
	er := stmt.Get(&id, usr)
	fmt.Printf("%T %v", id, id)
	return er, int32(id)

}

func (s *StoreDB) UserAuth(email string) (u storage.Users) {
	var user storage.Users
	err := s.Db.Get(&user, getUser, email)

	if err != nil {
		log.Printf("error for getting row: %v", err)
	}

	return user
}

func (s *StoreDB) UniqueEmailUname(email string) (emailId int32, unameID int32) {
	err := s.Db.Get(&emailId, getIdForEmail, email)
	if err != nil {
		log.Printf("error for getting row in finding uniqueness(email): %v", err)
	}

	err = s.Db.Get(&unameID, getIdForUname, email)
	if err != nil {
		log.Printf("error for getting row in finding uniqueness(username): %v", err)

	}

	fmt.Printf("uniqueness check from db: %T: %v;;;;; %T: %v", emailId, emailId, unameID, unameID)
	return emailId, unameID

}
