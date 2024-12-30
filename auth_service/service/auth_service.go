package service

import (
	"context"
	"fmt"
	"project_chat_app/auth_service/helper"
	"project_chat_app/auth_service/model"
	pb "project_chat_app/auth_service/proto"
	"project_chat_app/auth_service/repository"

	"go.uber.org/zap"
)

type AuthService struct {
	Repo  repository.AuthRepository
	Email EmailService
	Log   *zap.Logger
	pb.UnimplementedAuthServiceServer
}

// Login function handles user login request
func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	// Find user by email
	user, err := s.Repo.FindByEmail(req.Email)
	if err != nil {
		s.Log.Error("Failed to find user by email", zap.Error(err))
		return &pb.AuthResponse{Status: false, Message: "User not found"}, nil
	}

	// Generate a random OTP
	otp := helper.GenerateOTP()

	// Prepare email data
	emailData := map[string]interface{}{
		"FullName": fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		"OTP":      otp,
	}

	// Send OTP via email
	subject := "Your Login OTP"
	_, err = s.Email.Send(user.Email, subject, "otp_template", emailData)
	if err != nil {
		s.Log.Error("Failed to send OTP email", zap.Error(err))
		return &pb.AuthResponse{Status: false, Message: "Failed to send OTP email"}, nil
	}

	// Return success response
	return &pb.AuthResponse{Status: true, Message: "Login is successful. OTP sent to your email."}, nil
}

// Register function handles user registration request
func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	userInput := model.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	err := s.Repo.Create(&userInput)
	if err != nil {
		s.Log.Error("Failed to register user", zap.Error(err))
		return &pb.AuthResponse{Status: false, Message: "Failed to register user"}, nil
	}
	return &pb.AuthResponse{Status: true, Message: "Register is Success"}, nil
}

func (s *AuthService) VerifyOTP(ctx context.Context, req *pb.OTPRequest) (*pb.OTPResponse, error) {
	// Implement OTP verification logic here
	fmt.Printf("OTP request: %v\n", req)
	return &pb.OTPResponse{Status: true, Message: "OTP verified", UserEmail: "example@example.com", Token: "asdfghjkl"}, nil
}
