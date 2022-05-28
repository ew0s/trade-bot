package entities

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UID          string
	Name         string
	Username     string
	PasswordHash string
}

func (e *User) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(e.PasswordHash), []byte(password))
	return err == nil
}
