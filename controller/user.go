package controller

import (
	"AuthInGo/dto"
	"AuthInGo/service"
	"AuthInGo/utils"
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
	var payload dto.RegisterUserRequestDTO
	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", jsonErr)
		return
	}
	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error", validationErr)
		return
	}
	data, e := c.userService.CreateUser(&payload)
	if e != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "User creation failed", e)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Created user", data)
}
func (c *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginUserRequestDTO
	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request body", jsonErr)
		return
	}
	fmt.Println("Payload received:", payload)

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error", validationErr)
		return
	}
	token, err := c.userService.LoginUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Login failed", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Login successful", token)
}
