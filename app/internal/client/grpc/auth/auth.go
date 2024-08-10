package auth

import (
	"context"
	"time"

	"gophKeeper/internal/client/grpc"
	"gophKeeper/internal/lib/auth/keyring"
	pb "gophKeeper/protobuf/V1/users"
)

const defaultTimeout = 10 * time.Second

type AuthService struct {
	client pb.UserServiceClient
}

func NewAuthService() (*AuthService, error) {
	conn, err := grpc.GetConnUnauthenticated()
	if err != nil {
		return nil, err
	}
	return &AuthService{
		client: pb.NewUserServiceClient(conn),
	}, nil
}

func (s *AuthService) RegisterUser(login, password string, platform pb.Platform) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	req := &pb.UserRegistrationRequest{
		Login:    login,
		Password: password,
		Platform: platform,
	}

	_, err := s.client.Registration(ctx, req)
	return err
}

func (s *AuthService) LoginUser(login, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	req := &pb.LoginRequest{
		Login:    login,
		Password: password,
	}

	resp, err := s.client.Login(ctx, req)
	if err != nil {
		return "", err
	}

	if err := keyring.SaveToken(resp.Token); err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (s *AuthService) Logout() error {
	return keyring.DeleteToken()
}
