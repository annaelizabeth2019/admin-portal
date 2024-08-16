package router

import (
	"admin-portal/controller"
	"admin-portal/pkg/logger"
	"admin-portal/services/middleware"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	// Write gin access log to file
	f, err := os.OpenFile("cat_api.access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Errorf("Failed to create access log file: %v", err)
	} else {
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// Set default middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Set custom middlewares
	r.Use(middleware.CORS())
	r.Use(middleware.Security())
	r.Use(middleware.GetUserID())

	loginController := controller.NewLogin(db)

	loginRouter := r.Group("account")
	{
		loginRouter.GET("/login", loginController.Get)
		loginRouter.POST("/login", loginController.Submit)
		loginRouter.PATCH("/logout", loginController.Logout)
	}

	return r
}
