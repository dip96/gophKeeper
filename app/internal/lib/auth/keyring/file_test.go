package keyring

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveAndNewTokenAuth(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "gophkeeper_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	token := "test_token"
	err = SaveToken(token)
	assert.NoError(t, err)

	tokenPath := getTokenFilePath()
	savedToken, err := os.ReadFile(tokenPath)
	assert.NoError(t, err)
	assert.Equal(t, token, string(savedToken))

	tokenAuth, err := NewTokenAuth()
	assert.NoError(t, err)
	assert.Equal(t, token, tokenAuth.token)
}

func TestDeleteToken(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "gophkeeper_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	token := "test_token"
	err = SaveToken(token)
	require.NoError(t, err)

	err = DeleteToken()
	assert.NoError(t, err)

	_, err = os.Stat(getTokenFilePath())
	assert.True(t, os.IsNotExist(err))
}

func TestTokenAuth_GetRequestMetadata(t *testing.T) {
	token := "test_token"
	tokenAuth := &TokenAuth{token: token}

	metadata, err := tokenAuth.GetRequestMetadata(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"authorization": "Bearer " + token}, metadata)
}

func TestTokenAuth_RequireTransportSecurity(t *testing.T) {
	tokenAuth := &TokenAuth{}
	assert.False(t, tokenAuth.RequireTransportSecurity())
}

func TestGetTokenFilePath(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	require.NoError(t, err)

	expectedPath := filepath.Join(homeDir, "."+serviceName, fileName)
	actualPath := getTokenFilePath()

	assert.Equal(t, expectedPath, actualPath)
}
