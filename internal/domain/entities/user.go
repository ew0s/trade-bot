package entities

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UID          string `db:"uid" goqu:"skipinsert,skipupdate"`
	Name         string `db:"name"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

func (e *User) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(e.PasswordHash), []byte(password))
	return err == nil
}
