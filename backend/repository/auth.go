package repository

import (
	"admin-portal/model"

	"gorm.io/gorm"
)

type Auth interface {
	SelectByEmail(email string) (result model.Auth, err error)
	SetPassword(hash string, userID int64) (err error)
}

type auth struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) Auth {
	return &auth{
		db: db,
	}
}

func (r *auth) SelectByEmail(email string) (result model.Auth, err error) {
	return result, err
}

func (r *auth) SetPassword(hash string, userID int64) (err error) {
	tx := r.db.Table("admin_users").
		Where("user_id = ?", userID).
		Update("password", hash)
	return tx.Error
}
