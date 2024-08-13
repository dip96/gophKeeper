package config

import (
	"encoding/json"
	"errors"
	"flag"
	ErrorConfig "gophKeeper/internal/errors/config"
	"os"
	"sync"
)

// ConfigClient - интерфейс для конфига клиента
type ConfigClient interface {
	GetGrpcAddress() string
	GetSslCert() string
}

// client представляет конфигурацию клиента.
type client struct {
	// GrpcAddress - адрес gRPC сервера
	GrpcAddress string `json:"grpc_address"`
	// SslCert - путь к SSL сертификату
	SslCert string `json:"ssl_cert"`
}

// Реализация интерфейса ConfigClient для структуры client
func (c *client) GetGrpcAddress() string {
	return c.GrpcAddress
}

func (c *client) GetSslCert() string {
	return c.SslCert
}

var (
	instanceClient ConfigClient
	initOnceClient sync.Once
)

// MustLoadClient - загружает и инициализирует конфигурацию клиента
func MustLoadClient() ConfigClient {
	initOnceClient.Do(func() {
		var err error
		instanceClient, err = initClientConfig()

		if err != nil {
			var configErr *ErrorConfig.ErrorConfig

			if errors.As(err, &configErr) {
				panic(err)
			}

			panic(ErrorConfig.New("failed to initialize client config", err))
		}
	})

	return instanceClient
}

// initClientConfig инициализирует конфигурацию клиента
func initClientConfig() (ConfigClient, error) {
	var cfg client

	if err := parseClientFlags(&cfg); err != nil {
		return nil, ErrorConfig.New("error in parsing client flags", err)
	}

	if err := overrideClientFromEnv(&cfg); err != nil {
		return nil, ErrorConfig.New("error when overwriting client env variables", err)
	}

	// Чтение конфигурации из файла, если указан путь
	if configPath := os.Getenv("CLIENT_CONFIG_PATH"); configPath != "" {
		err := readConfigFileClient(configPath, &cfg)
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

func parseClientFlags(cfg *client) error {
	flag.StringVar(&cfg.GrpcAddress, "grpc", "localhost:3201", "gRPC server address")
	//flag.StringVar(&cfg.GrpcAddress, "grpc", "172.20.255.21:3201", "gRPC server address")
	flag.StringVar(&cfg.SslCert, "ssl_cert", "./ssl/cert.pem", "path to SSL certificate")
	flag.Parse()
	return nil
}

func overrideClientFromEnv(cfg *client) error {
	if envGrpcAddress := os.Getenv("GRPC_ADDRESS"); envGrpcAddress != "" {
		cfg.GrpcAddress = envGrpcAddress
	}
	if envSslCert := os.Getenv("SSL_CERT"); envSslCert != "" {
		cfg.SslCert = envSslCert
	}
	return nil
}

func readConfigFileClient(path string, cfg *client) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		return ErrorConfig.New("error parsing client configuration file", err)
	}

	return nil
}
