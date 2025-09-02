package controller

import (
	"AuthInGo/dto"
	"AuthInGo/service"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"
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
	payload := r.Context().Value("payload").(dto.RegisterUserRequestDTO)
	data, e := c.userService.CreateUser(&payload)
	if e != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "User creation failed", e)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Created user", data)
}
func (c *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginUserRequestDTO
	payload = r.Context().Value("payload").(dto.LoginUserRequestDTO)
	token, err := c.userService.LoginUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Login failed", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Login successful", token)
}
func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching user by ID in UserController")

	userId := r.URL.Query().Get("id")
	fmt.Println("User ID:", userId)
	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("missing user ID"))
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid user ID format", err)
		return
	}
	user, err := uc.userService.GetUserById(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}
	if user == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("user with ID %d not found", id))
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
	fmt.Println("User fetched successfully:", user)
}
