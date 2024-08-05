package grpc

import (
	"google.golang.org/grpc"
	"gophKeeper/internal/lib/auth/keyring"
)

func GetConn() (*grpc.ClientConn, error) {
	tokenAuth, err := keyring.NewTokenAuth()
	if err != nil {
		return nil, err
	}

	return grpc.Dial(
		"172.20.255.21:3201",
		//"localhost:3201",
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(tokenAuth),
	)
	//return grpc.Dial("172.20.255.21:3201", grpc.WithInsecure())
}
