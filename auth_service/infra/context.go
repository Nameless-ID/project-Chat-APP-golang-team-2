package infra

import (
	"project_chat_app/auth_service/config"
	"project_chat_app/auth_service/database"
	"project_chat_app/auth_service/infra/jwt"
	"project_chat_app/auth_service/log"
	"project_chat_app/auth_service/repository"
	"project_chat_app/auth_service/service"
)

type ServiceContext struct {
	Service *service.Service
	Cacher  database.Cacher
	JWT     jwt.JWT
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

	rdb := database.NewCacher(appConfig, 60*60)

	jwtLib := jwt.NewJWT(appConfig.PrivateKey, appConfig.PublicKey, logger)

	repo := repository.NewRepository(db, logger)
	return &ServiceContext{
		Service: service.NewService(*repo, appConfig, logger, rdb, jwtLib),
	}, nil
}
