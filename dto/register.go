package dto

type RegisterUserRequestDTO struct {
	UserName string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
