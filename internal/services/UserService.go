package services

import (
	"BeatsPro/internal/models/User"
	"BeatsPro/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) CreateUser(user *User.User) (int, error) {
	return userService.userRepository.CreateUser(user)
}

func (userService *UserService) UpdateUser(user *User.User) error {
	return userService.userRepository.UpdateUser(user)
}

func (userService *UserService) GetById(id int) (*User.User, error) {
	return userService.userRepository.GetById(id)
}
