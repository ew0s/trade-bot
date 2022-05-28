package models

type User struct {
	UID          string `db:"uid" goqu:"skipinsert,skipupdate"`
	Name         string `db:"name"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}
