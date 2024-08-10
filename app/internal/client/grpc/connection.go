package grpc

import (
	"google.golang.org/grpc"
	"gophKeeper/internal/config"
	"gophKeeper/internal/lib/auth/keyring"
)

func GetConnUnauthenticated() (*grpc.ClientConn, error) {
	return grpc.Dial(
		getAddr(),
		grpc.WithInsecure(),
	)
}

func GetConnAuthenticated() (*grpc.ClientConn, error) {
	tokenAuth, err := keyring.NewTokenAuth()
	if err != nil {
		return nil, err
	}

	return grpc.Dial(
		getAddr(),
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(tokenAuth),
	)
}

func getAddr() string {
	clientConfig := config.MustLoadClient()
	grpcAddress := clientConfig.GetGrpcAddress()
	return grpcAddress
}
