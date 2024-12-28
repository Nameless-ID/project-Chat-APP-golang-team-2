package main

import (
	// "log"
	// "user-service/controller"
	// "user-service/routes"
	// "github.com/gin-gonic/gin"
	"user-service/config"
	"user-service/grpc"
	"user-service/repository"
	"user-service/service"

)

func main() {

	db := config.SetupDatabase()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	// userController := controller.NewUserController(userService)

	// router := gin.Default()

	// routes.SetupUserRoutes(router, userController)

	// if err := router.Run(":8080"); err != nil {
	// 	log.Fatalf("Could not start server: %v", err)
	// }
	grpc.StartGRPCServer(userService)
}

