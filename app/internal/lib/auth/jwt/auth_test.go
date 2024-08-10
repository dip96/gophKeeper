package auth

import (
	"testing"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	secret := "test_secret"
	Init(secret)
	assert.NotNil(t, tokenAuth, "tokenAuth should not be nil after Init")
}

func TestGenerateToken(t *testing.T) {
	secret := "test_secret"
	Init(secret)

	userID := int64(123)
	token, err := GenerateToken(userID)

	assert.NoError(t, err, "GenerateToken should not return an error")
	assert.NotEmpty(t, token, "Generated token should not be empty")
}

func TestVerifyToken(t *testing.T) {
	secret := "test_secret"
	Init(secret)

	userID := int64(123)
	token, _ := GenerateToken(userID)

	claims, err := VerifyToken(token)

	assert.NoError(t, err, "VerifyToken should not return an error for a valid token")
	assert.NotNil(t, claims, "Claims should not be nil for a valid token")
	assert.Equal(t, userID, int64(claims["user_id"].(float64)), "UserID in claims should match the original userID")
}

func TestVerifyExpiredToken(t *testing.T) {
	secret := "test_secret"
	Init(secret)

	userID := int64(123)
	claims := map[string]interface{}{
		"user_id": userID,
		"exp":     time.Now().Add(-1 * time.Hour).Unix(),
	}
	_, token, _ := tokenAuth.Encode(claims)

	_, err := VerifyToken(token)

	assert.Error(t, err, "VerifyToken should return an error for an expired token")
	assert.Equal(t, jwtauth.ErrExpired, err, "Error should be jwtauth.ErrExpired for an expired token")
}

func TestVerifyInvalidToken(t *testing.T) {
	secret := "test_secret"
	Init(secret)

	invalidToken := "invalid.token.string"

	_, err := VerifyToken(invalidToken)

	assert.Error(t, err, "VerifyToken should return an error for an invalid token")
}
