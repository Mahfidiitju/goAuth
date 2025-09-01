package service

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetuUserById() error
	LoginUser() error
	CreateUser() error
}
type userServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: _userRepository,
	}
}

func (s *userServiceImpl) CreateUser() error {
	fmt.Println("create called")
	password := "password123"
	hassedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	fmt.Println(hassedPassword)
	s.userRepository.Create("testuser3", "test@test3.com", hassedPassword)
	return nil
}
func (s *userServiceImpl) GetuUserById() error {

	s.userRepository.GetByID()
	return nil
}
func (s *userServiceImpl) LoginUser() error {
	user, err := s.userRepository.GetUserByEmail("mahfidantor98@gmail.com")
	if err != nil {
		return err
	}
	fmt.Println(user)
	response := utils.CheckPasswordHash("password123", user.Password)
	if !response {
		fmt.Println("invalid password")
	}
	token, err := utils.CreateJWTToken(user.Id, user.Email)
	if err != nil {
		fmt.Println("invalid token", err)
	}
	fmt.Println("JWT Token:", token)
	return nil
}
