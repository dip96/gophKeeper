package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/service/login_data"
	pb "gophKeeper/protobuf/V1/login_data"
	"strconv"
)

type LoginDataService struct {
	pb.UnimplementedLoginDataServiceServer
	loginDataService *login_data.LoginDataService
}

func NewLoginDataService(loginDataService *login_data.LoginDataService) *LoginDataService {
	return &LoginDataService{loginDataService: loginDataService}
}

func (s *LoginDataService) SaveLoginData(ctx context.Context, req *pb.LoginDataRequest) (*pb.LoginDataResponse, error) {
	if req.Login == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login and password are required")
	}

	success, err := s.loginDataService.SaveLoginData(ctx, req.Login, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.LoginDataResponse{Success: success}, nil
}

func (s *LoginDataService) GetLoginData(ctx context.Context, req *pb.GetLoginDataRequest) (*pb.GetLoginDataResponse, error) {
	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	loginData, err := s.loginDataService.GetLoginData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetLoginDataResponse{
		Login:    loginData.Username,
		Password: string(loginData.Password),
		Id:       strconv.Itoa(int(loginData.ID)),
	}, nil
}

func (s *LoginDataService) GetAllLoginData(ctx context.Context, req *pb.GetAllLoginDataRequest) (*pb.GetAllLoginDataResponse, error) {
	responseItems, err := s.loginDataService.GetAllLoginData(ctx, 1, 10)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetAllLoginDataResponse{
		Items: responseItems,
	}, nil
}

func (s *LoginDataService) EditLoginData(ctx context.Context, req *pb.EditLoginDataRequest) (*pb.LoginDataResponse, error) {
	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.loginDataService.EditLoginData(ctx, entryID, req.Login, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.LoginDataResponse{Success: success}, nil
}

func (s *LoginDataService) DeleteLoginData(ctx context.Context, req *pb.DeleteLoginDataRequest) (*pb.LoginDataResponse, error) {
	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.loginDataService.DeleteLoginData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.LoginDataResponse{Success: success}, nil
}
