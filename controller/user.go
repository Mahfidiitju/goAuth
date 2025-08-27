package controller

import (
	"AuthInGo/service"
	"fmt"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(_userService service.UserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registering user")
	c.userService.CreateUser()
	w.Write([]byte("User registered successfully"))
}
func (c *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registering user")
	c.userService.LoginUser()
	w.Write([]byte("User registered successfully"))
}
