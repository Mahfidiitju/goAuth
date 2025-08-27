package service

import db "AuthInGo/db/repositories"

type UserService interface {
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
	s.userRepository.GetByID()
	return nil
}
