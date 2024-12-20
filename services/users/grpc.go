package main

import (
	"log"
	"net"

	"github.com/pdda-project/backend/services/users/pb"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	lis net.Listener
	pb.UnimplementedUserServiceServer
}

func NewGRPCServer(host string) (*gRPCServer, error) {
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}
	return &gRPCServer{lis: lis}, nil
}

func (server *gRPCServer) Run() {
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, server)
	log.Printf("server running on %s", server.lis.Addr())
	if err := s.Serve(server.lis); err != nil {
		log.Fatal(err.Error())
	}
}

func (server *gRPCServer) Close() {
	server.lis.Close()
}
