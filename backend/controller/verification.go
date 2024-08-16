package controller

import (
	"gorm.io/gorm"
)

type VerificationController struct {
	DB *gorm.DB
}
