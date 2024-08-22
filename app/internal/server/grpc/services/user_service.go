package services

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/service/user"
	pb "gophKeeper/protobuf/V1/users"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userService *user.UserService
}

func NewUserService(userService *user.UserService) *UserService {
	return &UserService{userService: userService}
}

func (s *UserService) Registration(ctx context.Context, req *pb.UserRegistrationRequest) (*pb.UserRegistrationResponse, error) {
	if req.Login == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login and password are required")
	}

	success, err := s.userService.RegisterUser(ctx, req.Login, req.Password)
	if err != nil {
		switch err {
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.UserRegistrationResponse{
		Success: success,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Login == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login and password are required")
	}

	token, err := s.userService.AuthenticateUser(ctx, req.Login, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, status.Error(codes.NotFound, "user not found")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.LoginResponse{
		Success: true,
		Token:   token,
	}, nil
}
