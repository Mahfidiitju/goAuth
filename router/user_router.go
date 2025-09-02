package router

import (
	"AuthInGo/controller"
	"AuthInGo/middlewares"

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
	r.With(middlewares.UserCreateRequestValidator).Post("/signup", ur.userController.RegisterUser)
	r.With(middlewares.UserLoginRequestValidator).Post("/login", ur.userController.LoginUser)
	r.With(middlewares.JWTAuthMiddleware).Get("/user", ur.userController.GetUserById)
	r.With(middlewares.JWTAuthMiddleware).Get("/user/all", ur.userController.GetAllUsers)
}
