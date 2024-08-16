package controller

import (
	"admin-portal/services/admin"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type loginController struct {
	DB           *gorm.DB
	adminService admin.Service
}

type LoginController interface {
	Get(c *gin.Context)
	Submit(c *gin.Context)
	Logout(c *gin.Context)
}

func NewLogin(db *gorm.DB) LoginController {
	return &loginController{
		DB: db,
	}
}

func (ctlr *loginController) Get(c *gin.Context) {
	c.HTML(200, "login.tmpl", nil)
}
func (ctlr *loginController) Submit(c *gin.Context) {
	form := LoginForm{}
	err := c.MustBindWith(&form, binding.Form)
	if err != nil {
		// show an error page
		return
	}
	err = validate.Struct(form)
	if err != nil {
		// do something to indicate there was an error
		return
	}

	user, err := ctlr.adminService.Authenticate(form.Email, form.Password)
	if err != nil {
		// show an error page
		return
	}
	c.JSON(200, user)
}
func (ctlr *loginController) Logout(c *gin.Context) {
	c.Status(200)
}
