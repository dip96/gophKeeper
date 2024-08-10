package auth

import (
	"sync"
)

var (
	authService *AuthService
	once        sync.Once
)

func GetAuthService() (*AuthService, error) {
	var err error
	once.Do(func() {
		authService, err = NewAuthService()
	})
	return authService, err
}
