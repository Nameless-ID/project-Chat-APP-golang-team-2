package main

import (
	"flag"
	"log"
	"net"
	"project_chat_app/auth_service/infra"

	pb "project_chat_app/auth_service/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal(err)
	}


	var listener net.Listener
	listener, err = net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &ctx.Service.Auth)
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
	log.Println("Server running on port 50051")
}

