package router

import (
	"AuthInGo/controller"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controller.UserController
}

func NewUserRouter(_userController *controller.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Post("/signup", ur.userController.RegisterUser)
	r.Post("/login", ur.userController.LoginUser)
}
