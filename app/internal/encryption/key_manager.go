package encryption

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
)

// TODO add config
const keyFilePath = "./keys/keyfile"

func SaveKeyToFile(key []byte) error {
	encodedKey := base64.StdEncoding.EncodeToString(key)
	return ioutil.WriteFile(keyFilePath, []byte(encodedKey), 0644)
}

func LoadKeyFromFile() ([]byte, error) {
	if _, err := os.Stat(keyFilePath); os.IsNotExist(err) {
		return nil, errors.New("key file does not exist")
	}

	data, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(string(data))
}
