package controller

import (
	"gorm.io/gorm"
)

type UsersController struct {
	DB *gorm.DB
}
