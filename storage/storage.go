package storage

import (
	"time"
)

type Users struct {
	ID        int32     `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	IsActive  bool      `db:"is_active"`
	IsAdmin   bool      `db:"is_admin"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type SignupUser struct {
	FirstName string `schema:"firstname" db:"first_name"`
	LastName  string `schema:"lastname" db:"last_name"`
	Username  string `schema:"username" db:"username"`
	Email     string `schema:"email" db:"email"`
	Password  string `schema:"password" db:"password"`
}
