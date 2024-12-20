package main

import (
	"context"

	"github.com/pdda-project/backend/services/users/pb"
)

func (s *gRPCServer) ChangeUserEmail(ctx context.Context, req *pb.ChangeUserEmailRequest) (*pb.ChangeUserEmailResponse, error) {
	panic("unimplemented")
}

func (s *gRPCServer) ChangeUserPassword(ctx context.Context, req *pb.ChangeUserPasswordRequest) (*pb.ChangeUserPasswordResponse, error) {
	panic("unimplemented")
}

func (s *gRPCServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	panic("unimplemented")
}

func (s *gRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	panic("unimplemented")
}

func (s *gRPCServer) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	panic("unimplemented")
}

func (s *gRPCServer) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	panic("unimplemented")
}

func (s *gRPCServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	panic("unimplemented")
}
