package service

import "project_chat_app/auth_service/repository"

type Service struct {
	Auth AuthService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Auth: AuthService{Repo: repo.Auth},
	}
}
