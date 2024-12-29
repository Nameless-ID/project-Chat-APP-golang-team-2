package service

import (
	"context"
	"fmt"
	"project_chat_app/auth_service/model"
	pb "project_chat_app/auth_service/proto"
	"project_chat_app/auth_service/repository"
)

type AuthService struct {
	Repo repository.AuthRepository
	pb.UnimplementedAuthServiceServer
}

// Login function handles user login request
func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	user, err := s.Repo.FindByEmail(req.Email)
	if err != nil {
		return &pb.AuthResponse{Status: false, Message: "User not found"}, nil
	}
	fmt.Printf("user data: %v\n", user)
	return &pb.AuthResponse{Status: true, Message: "Login is Success"}, nil
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
		return &pb.AuthResponse{Status: false, Message: "Failed to register user"}, nil
	}
	return &pb.AuthResponse{Status: true, Message: "Register is Success"}, nil
}

func (s *AuthService) VerifyOTP(ctx context.Context, req *pb.OTPRequest) (*pb.OTPResponse, error) {
	// Implement OTP verification logic here
	fmt.Printf("OTP request: %v\n", req)
	return &pb.OTPResponse{Status: true, Message: "OTP verified", UserEmail: "example@example.com", Token: "asdfghjkl"}, nil
}
