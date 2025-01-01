package grpc

import (
	"context"
	"log"
	"net"
	"strconv"
	"user-service/models"
	"user-service/proto"
	"user-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	userService service.UserService
	proto.UnimplementedUserServiceServer
}

func NewServer(userService service.UserService) *server {
	return &server{userService: userService}
}

func (s *server) GetAllUsers(ctx context.Context, req *proto.Empty) (*proto.UsersList, error) {
	users, err := s.userService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var userResponses []*proto.User
	for _, user := range users {
		userResponses = append(userResponses, &proto.User{
			Id:        int32(user.ID),
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			IsOnline:  user.IsOnline,
		})

	}
	return &proto.UsersList{Users: userResponses}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {

	if req.FirstName == "" || req.LastName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "First name or last name cannot be empty")
	}

	existingUser, err := s.userService.GetUserInfo(strconv.Itoa(int(req.Id)))
	if err != nil {
		if err.Error() == "user not found" {
			return nil, status.Errorf(codes.NotFound, "User with ID %d not found", req.Id)
		}
		log.Printf("Error fetching user info: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to fetch user info")
	}

	user := &models.User{
		ID:        existingUser.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := s.userService.UpdateUser(user); err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to update user: %v", err)
	}

	log.Printf("User with ID %d updated successfully", req.Id)
	return &proto.UpdateUserResponse{Message: "Update success"}, nil
}

func StartGRPCServer(userService service.UserService) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, NewServer(userService))
	reflection.Register(grpcServer)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
