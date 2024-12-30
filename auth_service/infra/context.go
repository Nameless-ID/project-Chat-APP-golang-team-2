package infra

import (
	"project_chat_app/auth_service/config"
	"project_chat_app/auth_service/database"
	"project_chat_app/auth_service/log"
	"project_chat_app/auth_service/repository"
	"project_chat_app/auth_service/service"
)

type ServiceContext struct {
	Service *service.Service
}

func NewServiceContext() (*ServiceContext, error) {
	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	appConfig, err := config.LoadConfig()
	if err != nil {
		return handlerError(err)
	}

	db, err := database.ConnectDB(appConfig)
	if err != nil {
		return handlerError(err)
	}

	logger, err := log.InitZapLogger(appConfig)
	if err != nil {
		return handlerError(err)
	}

	repo := repository.NewRepository(db, logger)
	return &ServiceContext{
		Service: service.NewService(*repo, appConfig, logger),
	}, nil
}
