package keyring

import (
	"context"
	"errors"
	"github.com/zalando/go-keyring"
	"google.golang.org/grpc"
)

const (
	serviceName = "GophKeeper"
	tokenKey    = "AuthToken"
)

// SaveToken сохраняет токен в системном keyring
func SaveToken(token string) error {
	return keyring.Set(serviceName, tokenKey, token)
}

// LoadToken загружает токен из системного keyring
func LoadToken() (string, error) {
	return keyring.Get(serviceName, tokenKey)
}

// DeleteToken удаляет токен из системного keyring
func DeleteToken() error {
	return keyring.Delete(serviceName, tokenKey)
}

// TokenAuth реализует grpc.PerRPCCredentials для аутентификации запросов
type TokenAuth struct {
	token string
}

func (t TokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.token,
	}, nil
}

func (TokenAuth) RequireTransportSecurity() bool {
	return false // Измените на true, если используете TLS
}

// GetAuthenticatedClient возвращает gRPC клиент с аутентификацией
func GetAuthenticatedClient(target string) (*grpc.ClientConn, error) {
	token, err := LoadToken()
	if err != nil {
		return nil, errors.New("не удалось загрузить токен аутентификации")
	}

	return grpc.Dial(
		target,
		grpc.WithInsecure(), // Замените на grpc.WithTransportCredentials(...) для TLS
		grpc.WithPerRPCCredentials(TokenAuth{token: token}),
	)
}
