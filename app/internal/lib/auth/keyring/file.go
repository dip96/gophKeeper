package keyring

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	serviceName = "GophKeeper"
	fileName    = "auth_token.txt"
)

func getTokenFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "."+serviceName, fileName)
}

func SaveToken(token string) error {
	tokenPath := getTokenFilePath()
	err := os.MkdirAll(filepath.Dir(tokenPath), 0700)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(tokenPath, []byte(token), 0600)
}

func NewTokenAuth() (*TokenAuth, error) {
	tokenPath := getTokenFilePath()
	tokenBytes, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}
	return &TokenAuth{token: string(tokenBytes)}, nil
}

func DeleteToken() error {
	tokenPath := getTokenFilePath()
	return os.Remove(tokenPath)
}

// реализует grpc.PerRPCCredentials для аутентификации запросов
type TokenAuth struct {
	token string
}

func (t TokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.token,
	}, nil
}

func (TokenAuth) RequireTransportSecurity() bool {
	return false
}
