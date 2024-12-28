package service

import (
	"user-service/models"
	"user-service/repository"
)

type UserService interface {
    GetUserInfo(userID string) (*models.User, error)
    GetOnlineUsers() ([]models.User, error)
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo}
}

func (s *userService) GetUserInfo(userID string) (*models.User, error) {
    return s.repo.GetUserByID(userID)
}

func (s *userService) GetOnlineUsers() ([]models.User, error) {
    return s.repo.GetAllUsers()
}
