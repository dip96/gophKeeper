package auth

import (
	"context"
	"gophKeeper/internal/client/grpc"
	auth "gophKeeper/internal/lib/auth/keyring"
	pb "gophKeeper/protobuf/V1/users"
)

func RegisterUser(login, password string, platform pb.Platform) error {
	conn, err := grpc.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	req := &pb.UserRegistrationRequest{
		Login:    login,
		Password: password,
		Platform: platform,
	}

	_, err = client.Registration(context.Background(), req)
	return err
}

func LoginUser(login, password string) (string, error) {
	conn, err := grpc.GetConn()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	req := &pb.LoginRequest{
		Login:    login,
		Password: password,
	}

	resp, err := client.Login(context.Background(), req)
	if err != nil {
		return "", err
	}

	if err := auth.SaveToken(resp.Token); err != nil {
		return "", err
	}

	return resp.Token, nil
}

func Logout() error {
	if err := auth.DeleteToken(); err != nil {
		return err
	}

	return nil
}
