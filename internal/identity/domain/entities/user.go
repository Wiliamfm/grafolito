package entities

import (
	"grafolito/backend/internal/shared/domain/entities"
)

type User struct {
	Id       string
	Username string
	Email    string
	password string
	Name     string
	LastName string
	*entities.Audit
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) Password() string {
	return u.password
}
