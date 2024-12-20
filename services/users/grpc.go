package main

import (
	"log"
	"net"

	"github.com/pdda-project/backend/services/users/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type gRPCServer struct {
	lis net.Listener
	db  *gorm.DB
	pb.UnimplementedUserServiceServer
}

func NewGRPCServer(host string, db *gorm.DB) (*gRPCServer, error) {
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}
	return &gRPCServer{lis: lis, db: db}, nil
}

func (server *gRPCServer) Run() {
	defer server.lis.Close()
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, server)
	log.Printf("server running on %s", server.lis.Addr())
	if err := s.Serve(server.lis); err != nil {
		log.Fatal(err.Error())
	}
}
