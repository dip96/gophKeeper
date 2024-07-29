package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

func GenerateRandomKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GenerateEncryptionKey(masterKey, salt []byte) []byte {
	hash := sha256.New()
	hash.Write(append(masterKey, salt...))
	return hash.Sum(nil)
}

func Encrypt(plainText, key []byte) (cipherText, iv []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	iv = make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	cipherText = make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText, iv, nil
}

func Decrypt(cipherText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText, nil
}

func EncodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func DecodeBase64(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}
