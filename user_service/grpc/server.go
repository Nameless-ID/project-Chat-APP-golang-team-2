package grpc

import (
	"context"
	"log"
	"net"
	"strconv"
	"user-service/proto"
	"user-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	userService service.UserService
	proto.UnimplementedUserServiceServer
}

func NewServer(userService service.UserService) *server {
	return &server{userService: userService}
}

func (s *server) GetUserInfo(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	user, err := s.userService.GetUserInfo(req.UserId)
	if err != nil {
		log.Printf("Error fetching user info for UserId %s: %v", req.UserId, err)
		return nil, err
	}
	return &proto.UserResponse{
		UserId:   strconv.Itoa(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
	}, nil
}

func (s *server) GetOnlineUsers(ctx context.Context, req *proto.Empty) (*proto.UsersList, error) {
	users, err := s.userService.GetOnlineUsers()
	if err != nil {
		return nil, err
	}
	var userResponses []*proto.UserResponse
	for _, user := range users {
		if user.IsActive {
			userResponses = append(userResponses, &proto.UserResponse{
				UserId:   strconv.Itoa(user.ID),
				Name:     user.Name,
				Email:    user.Email,
				IsActive: user.IsActive,
			})

		}
	}
	return &proto.UsersList{Users: userResponses}, nil
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
