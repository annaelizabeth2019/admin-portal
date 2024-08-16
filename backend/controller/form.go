package controller

type LoginForm struct {
	Email    string `form:"email" validate:"required"`
	Password string `form:"password" validate:"required"`
}
