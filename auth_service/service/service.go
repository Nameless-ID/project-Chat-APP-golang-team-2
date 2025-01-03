package service

import (
	"project_chat_app/auth_service/config"
	"project_chat_app/auth_service/database"
	"project_chat_app/auth_service/infra/jwt"
	"project_chat_app/auth_service/repository"

	"go.uber.org/zap"
)

type Service struct {
	Auth AuthService
}

func NewService(repo repository.Repository, config config.Config, log *zap.Logger, rdb database.Cacher, jwt jwt.JWT) *Service {
	return &Service{
		Auth: AuthService{Repo: repo.Auth, Email: NewEmailService(config.Email, log), Log: log, Cacher: rdb, Jwt: jwt},
	}
}
