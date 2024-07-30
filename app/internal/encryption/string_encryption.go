package encryption

import (
	"encoding/base64"
)

type StringEncryptor struct {
	encryptionKey []byte
}

func NewStringEncryptor(encryptionKey []byte) *StringEncryptor {
	return &StringEncryptor{encryptionKey: encryptionKey}
}

func (se *StringEncryptor) EncryptString(data string) (cipherText, iv string, err error) {
	cipherBytes, ivBytes, err := Encrypt([]byte(data), se.encryptionKey)
	if err != nil {
		return "", "", err
	}
	return base64.StdEncoding.EncodeToString(cipherBytes), base64.StdEncoding.EncodeToString(ivBytes), nil
}

func (se *StringEncryptor) DecryptString(cipherText, iv string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	plainText, err := Decrypt(cipherBytes, se.encryptionKey, ivBytes)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}
