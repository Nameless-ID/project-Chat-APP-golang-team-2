package main

import (
	"log"

	"api-gateway/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load konfigurasi
	initConfig()

	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	// gRPC connections
	userConn, err := grpc.Dial(viper.GetString("grpc.user_service"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to User Service: %v", err)
	}
	defer userConn.Close()
	userClient := pbUser.NewUserServiceClient(userConn)

	chatConn, err := grpc.Dial(viper.GetString("grpc.chat_service"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Chat Service: %v", err)
	}
	defer chatConn.Close()
	chatClient := pbChat.NewChatServiceClient(chatConn)

	// Router setup
	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(logger))

	// User routes
	router.POST("/user/register", func(c *gin.Context) { /* Handler */ })
	router.POST("/user/login", func(c *gin.Context) { /* Handler */ })

	// Chat routes
	chatRoutes := router.Group("/chat")
	chatRoutes.Use(middleware.AuthMiddleware(userClient))
	chatRoutes.Use(middleware.CacheMiddleware(redisClient))
	chatRoutes.POST("/send", func(c *gin.Context) { /* Handler */ })
	chatRoutes.GET("/messages", func(c *gin.Context) { /* Handler */ })

	router.Run(":" + viper.GetString("server.port"))
}
