package auth

import (
	"github.com/go-chi/jwtauth"
	"time"
)

var tokenAuth *jwtauth.JWTAuth

func Init(secret string) {
	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}

func GenerateToken(userID int64) (string, error) {
	claims := map[string]interface{}{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	_, tokenString, err := tokenAuth.Encode(claims)
	return tokenString, err
}

func VerifyToken(tokenString string) (map[string]interface{}, error) {
	token, err := tokenAuth.Decode(tokenString)
	if err != nil {
		return nil, err
	}
	if token.Expiration().Before(time.Now()) {
		return nil, jwtauth.ErrExpired
	}
	return token.PrivateClaims(), nil
}
