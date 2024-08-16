package model

type Auth struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password_hash"`
}
