package repository

import (
	"admin-portal/model"

	"gorm.io/gorm"
)

type User interface {
	SelectOne(userID int64) (model.User, error)
	GetRolesForUser(userID int64) ([]model.Role, error)
}

type user struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) User {
	return &user{
		db: db,
	}
}

func (repo *user) SelectOne(userID int64) (model.User, error) {
	return model.User{}, nil
}

func (repo *user) GetRolesForUser(userID int64) (roles []model.Role, err error) {
	return roles, err
}
