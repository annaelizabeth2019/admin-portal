package model

type User struct {
	ID    int64              `db:"id"`
	Name  string             `db:"name"`
	Email string             `db:"email"`
	Roles map[AdminRole]Role `db:"-"`
}

type AdminRole string

const (
	Owner     AdminRole = "owner"
	Viewer    AdminRole = "viewer"
	Moderator AdminRole = "moderator"
)

type Role struct {
	ID    int       `db:"id"`
	Title AdminRole `db:"title"`
}
