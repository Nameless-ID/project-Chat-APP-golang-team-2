package service

import (
	"user-service/models"
	"user-service/repository"
)

type UserService interface {
    GetUserInfo(userID string) (*models.User, error)
    GetAllUsers() ([]*models.User, error)
    UpdateUser(user *models.User) error
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

func (s *userService) GetAllUsers() ([]*models.User, error) {
    return s.repo.GetAllUsers()
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}
