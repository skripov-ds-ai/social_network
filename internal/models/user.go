package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserParams struct {
	Password   string
	FirstName  string
	SecondName string
	Birthdate  time.Time
	Biography  *string
	City       *string
}

type UpdateUserParams struct {
	ID           string
	PasswordHash *string
	FirstName    *string
	SecondName   *string
	Birthdate    *time.Time
	Biography    *string
	City         *string
}

type User struct {
	ID           uuid.UUID `db:"id"`
	PasswordHash string    `db:"password_hash"`
	FirstName    string    `db:"first_name"`
	SecondName   string    `db:"second_name"`
	Birthdate    time.Time `db:"birthdate"`
	Biography    *string   `db:"biography"`
	City         *string   `db:"city"`
}
