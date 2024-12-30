package service

import (
	"project_chat_app/auth_service/config"
	"project_chat_app/auth_service/repository"

	"go.uber.org/zap"
)

type Service struct {
	Auth AuthService
}

func NewService(repo repository.Repository, config config.Config, log *zap.Logger) *Service {
	return &Service{
		Auth: AuthService{Repo: repo.Auth, Email: NewEmailService(config.Email, log), Log: log},
	}
}
