package service

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"errors"
	"fmt"
)

type UserService interface {
	GetUserById(id int64) (*models.User, error)
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
	CreateUser(payload *dto.RegisterUserRequestDTO) (*models.User, error)
}
type userServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: _userRepository,
	}
}

func (s *userServiceImpl) CreateUser(payload *dto.RegisterUserRequestDTO) (*models.User, error) {
	password := payload.Password
	email := payload.Email
	userName := payload.UserName
	hassedPassword, err := utils.HashPassword(password)
	if err != nil {
		return &models.User{}, err
	}
	fmt.Println(hassedPassword)

	data, err := s.userRepository.Create(userName, email, hassedPassword)
	if err != nil {
		return &models.User{}, err
	}
	return data, nil
}
func (u *userServiceImpl) GetUserById(id int64) (*models.User, error) {
	fmt.Println("Fetching user in UserService")
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return nil, err
	}
	return user, nil
}

func (s *userServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	email := payload.Email
	password := payload.Password
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	fmt.Println("up", user.Password)
	fmt.Println("p", password)
	response := utils.CheckPasswordHash(password, user.Password)
	if !response {
		fmt.Println("invalid password")
		return "paassword not matched", errors.New("password not matched")
	}
	token, err := utils.CreateJWTToken(user.Id, user.Email)
	if err != nil {
		fmt.Println("invalid token", err)
		return "", err
	}
	fmt.Println("JWT Token:", token)
	return token, nil
}
