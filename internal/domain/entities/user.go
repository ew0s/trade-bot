package entities

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UID          string `db:"uid" goqu:"skipinsert,skipupdate"`
	Name         string `db:"name"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

func (e *User) GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("genrating hash from password: %w", err)
	}

	return string(bytes), nil
}

func (e *User) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(e.PasswordHash), []byte(password))
	return err == nil
}
