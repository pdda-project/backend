package main

import (
	"context"

	"log"

	"github.com/go-playground/validator/v10"
	"github.com/pdda-project/backend/services/users/pb"
	"github.com/pdda-project/backend/services/users/repo"
	"github.com/pdda-project/backend/utils"
	"github.com/pdda-project/backend/utils/errors"
	"gorm.io/gorm"
)

func (s *gRPCServer) ChangeUserEmail(
	ctx context.Context,
	req *pb.ChangeUserEmailRequest,
) (*pb.ChangeUserEmailResponse, error) {
	panic("unimplemented")
}

func (s *gRPCServer) ChangeUserPassword(
	ctx context.Context,
	req *pb.ChangeUserPasswordRequest,
) (*pb.ChangeUserPasswordResponse, error) {
	panic("unimplemented")
}

func (s *gRPCServer) DeleteUser(
	ctx context.Context,
	req *pb.DeleteUserRequest,
) (*pb.DeleteUserResponse, error) {
	panic("unimplemented")
}

func (s *gRPCServer) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {
	// Get user from db
	var user repo.User
	err := s.db.Where("uid = ?", req.Uid).First(&user).Error

	if err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewHTTPError(404, "user not found")
		} else {
			return nil, errors.InternalError
		}
	}

	// Turn into response
	return &pb.GetUserResponse{
		Uid:           user.Uid,
		DisplayName:   user.DisplayName,
		IsAdmin:       user.IsAdmin,
		IsExpert:      user.IsExpert,
		IsStaff:       user.IsStaff,
		RoleAtCompany: user.RoleAtCompany,
		Company:       user.Company,
		CompanyEmail:  user.CompanyEmail,
		CreatedAt:     user.CreatedAt.String(),
		PhotoUrl:      user.PhotoUrl,
	}, nil

}

func (s *gRPCServer) ListUser(
	ctx context.Context,
	req *pb.ListUserRequest,
) (*pb.ListUserResponse, error) {
	panic("unimplemented")
}

// Register User Serivice
func (s *gRPCServer) RegisterUser(
	ctx context.Context,
	req *pb.RegisterUserRequest,
) (*pb.RegisterUserResponse, error) {
	// Validator instance
	validate := validator.New()
	v := validate.Var

	// Validate request
	if err := v(req.Email, "required,email"); err != nil {
		log.Println(err.Error())
		return nil, errors.NewHTTPError(400, "insert a valid email")
	}
	if err := v(req.Password, "required,min=8"); err != nil {
		log.Println(err.Error())
		return nil, errors.NewHTTPError(400, "password minimum length is 8 characters")
	}

	// Get count email
	countUser, err := utils.CountRecord(s.db, &repo.User{}, "email", req.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.InternalError
	}
	// Check count email
	if countUser != 0 {
		return nil, errors.NewHTTPError(409, "email already exist")
	}

	// Write model
	var user repo.User
	user.Email = req.Email
	user.PasswordHash = req.Password

	// Create and write to db
	err = s.db.Create(&user).Error
	if err != nil {
		log.Println(err.Error())
		return nil, errors.InternalError
	}
	return &pb.RegisterUserResponse{Uid: user.Uid}, nil
}

// Update User Serivice
func (s *gRPCServer) UpdateUser(
	ctx context.Context,
	req *pb.UpdateUserRequest,
) (*pb.UpdateUserResponse, error) {
	panic("unimplemented")
}
