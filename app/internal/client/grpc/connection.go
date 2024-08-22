package grpc

import (
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gophKeeper/internal/config"
	"gophKeeper/internal/lib/auth/keyring"
)

func GetConnUnauthenticated() (*grpc.ClientConn, error) {
	clientConfig := config.MustLoadClient()

	creds, err := loadTLSCredentials(clientConfig)
	if err != nil {
		return nil, err
	}

	return grpc.Dial(
		getAddr(),
		grpc.WithTransportCredentials(creds),
	)
}

func GetConnAuthenticated() (*grpc.ClientConn, error) {
	tokenAuth, err := keyring.NewTokenAuth()
	if err != nil {
		return nil, err
	}

	clientConfig := config.MustLoadClient()
	creds, err := loadTLSCredentials(clientConfig)
	if err != nil {
		return nil, err
	}

	return grpc.Dial(
		getAddr(),
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(tokenAuth),
	)
}

func loadTLSCredentials(clientConfig config.ConfigClient) (credentials.TransportCredentials, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	return credentials.NewTLS(config), nil
}

func getAddr() string {
	clientConfig := config.MustLoadClient()
	return clientConfig.GetGrpcAddress()
}
