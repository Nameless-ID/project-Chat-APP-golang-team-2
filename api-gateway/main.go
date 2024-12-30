package main

import (
	"context"
	"log"
	"net/http"

	"project_chat_app/api-gateway/middleware"
	authpb "project_chat_app/auth-service"
	userpb "project_chat_app/user-service/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	authClient authpb.AuthServiceClient
	userClient userpb.UserServiceClient
)

func main() {
	// Inisialisasi koneksi gRPC ke Auth Service
	conn, err := grpc.Dial("103.127.132.149:4041", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}
	defer conn.Close()

	authClient = authpb.NewAuthServiceClient(conn)

	// Inisialisasi koneksi gRPC ke User Service
	userConn, err := grpc.Dial("103.127.132.149:4042", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to User Service: %v", err)
	}
	defer userConn.Close()
	userClient = userpb.NewUserServiceClient(userConn)

	router := gin.Default()

	// Routing untuk Auth
	router.POST("/auth/register", registerHandler)
	router.POST("/auth/login", loginHandler)
	router.POST("/auth/verify-otp", verifyOTPHandler)

	router.Use(middleware.Authentication())

	// Routing untuk User Service
	router.GET("/user/:id", getUserInfoHandler)
	router.GET("/users/online", getOnlineUsersHandler)

	log.Println("API Gateway running on port 4040...")
	router.Run(":4040")
}

// Handler untuk Register
func registerHandler(c *gin.Context) {
	var req authpb.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := authClient.Register(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": res.Status, "message": res.Message})
}

// Handler untuk Login
func loginHandler(c *gin.Context) {
	var req authpb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := authClient.Login(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": res.Status, "message": res.Message})
}

// Handler untuk Verify OTP
func verifyOTPHandler(c *gin.Context) {
	var req authpb.OTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := authClient.VerifyOTP(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     res.Status,
		"message":    res.Message,
		"user_email": res.UserEmail,
		"token":      res.Token,
	})
}

// Handler untuk mendapatkan informasi pengguna
func getUserInfoHandler(c *gin.Context) {
	userID := c.Param("id")

	res, err := userClient.GetUserInfo(context.Background(), &userpb.UserRequest{UserId: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":   res.UserId,
		"name":      res.Name,
		"email":     res.Email,
		"is_active": res.IsActive,
	})
}

// Handler untuk mendapatkan daftar pengguna online
func getOnlineUsersHandler(c *gin.Context) {
	res, err := userClient.GetOnlineUsers(context.Background(), &userpb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get online users"})
		return
	}

	users := []gin.H{}
	for _, user := range res.Users {
		users = append(users, gin.H{
			"user_id":   user.UserId,
			"name":      user.Name,
			"email":     user.Email,
			"is_active": user.IsActive,
		})
	}

	c.JSON(http.StatusOK, users)
}
