package config

import (
	"encoding/json"
	"errors"
	"flag"
	ErrorConfig "gophKeeper/internal/errors/config"
	"os"
	"sync"
)

// ConfigServer - интерфейс для конфига сервера
type ConfigServer interface {
	GetRunAddress() string
	GetDatabaseURI() string
	GetMigrationPath() string
	GetStand() string
	GetLoggerLevel() string
	GetGrpcAddress() string
	GetSslCert() string
	GetSslKey() string
}

// Реализация интерфейса ConfigInstance для структуры Server
func (c *server) GetRunAddress() string {
	return c.InitialConfig
}

func (c *server) GetDatabaseURI() string {
	return c.DatabaseDsn
}

func (c *server) GetMigrationPath() string {
	return c.MigrationPath
}

func (c *server) GetStand() string {
	return c.Stand
}
func (c *server) GetLoggerLevel() string {
	return c.LoggerLevel
}

func (c *server) GetGrpcAddress() string {
	return c.GrpcAddress
}

func (c *server) GetSslCert() string {
	return c.SslCert
}

func (c *server) GetSslKey() string {
	return c.SslKey
}

// server представляет конфигурацию сервера.
type server struct {
	// InitialConfig - путь до файла конфигурации
	InitialConfig string
	// DatabaseDsn - строка подключения к базе данных.
	DatabaseDsn string `json:"database_dsn"`
	// MigrationPath - путь к файлам миграций базы данных.
	MigrationPath string `json:"migration_path"`
	//Stand - наименование стенда
	Stand string `json:"stand"`
	//LoggerLevel - уровень логирования
	LoggerLevel string `json:"logger_level"`
	//GrpcAddress - адрес grpc
	GrpcAddress string `json:"grpc_address"`
	//SslCert - путь к SSL сертификату
	SslCert string `json:"ssl_cert"`
	// SslKey - путь к SSL ключу
	SslKey string `json:"ssl_key"`
}

var (
	instance ConfigServer
	initOnce sync.Once
)

// MustLoad - загружает и инициализирует конфигурацию сервера
// Функция обеспечивает однократную инициализацию конфигурации
func MustLoad() ConfigServer {
	// initConfig является синглтоном, что для конфига не является критичным, так как он инициализируется один раз
	// и не будет больше меняться
	initOnce.Do(func() {
		var err error
		instance, err = initConfig()

		if err != nil {
			var configErr *ErrorConfig.ErrorConfig

			if errors.As(err, &configErr) {
				panic(err)
			}

			panic(ErrorConfig.New("failed to initialize config", err))
		}
	})

	return instance
}

// initConfig инициализирует конфигурацию сервера на основе переданных флагов
// командной строки и переменных окружения.
func initConfig() (ConfigServer, error) {
	var cfg server

	if err := parseFlags(&cfg); err != nil {
		return nil, ErrorConfig.New("error in parsing", err)
	}

	if err := overrideFromEnv(&cfg); err != nil {
		return nil, ErrorConfig.New("error when overwriting env variables", err)
	}

	if cfg.InitialConfig != "" {
		err := readConfigFileServer(cfg.InitialConfig, &cfg)

		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

func parseFlags(cfg *server) error {
	flag.StringVar(&cfg.InitialConfig, "i", "", "path to init config")
	flag.StringVar(&cfg.DatabaseDsn, "d", "", "url for database")
	flag.StringVar(&cfg.MigrationPath, "m", "", "path to migration file")
	flag.StringVar(&cfg.Stand, "s", "", "current stand")
	flag.StringVar(&cfg.LoggerLevel, "l", "", "logger level")
	flag.StringVar(&cfg.GrpcAddress, "g", "", "grpc address")
	flag.StringVar(&cfg.SslCert, "ssl_cert", "", "path to SSL certificate")
	flag.StringVar(&cfg.SslKey, "ssl_key", "", "path to SSL key")

	flag.Parse()
	return nil
}

func overrideFromEnv(cfg *server) error {
	if envInitialConfig := os.Getenv("INITIAL_CONFIG"); envInitialConfig != "" {
		cfg.InitialConfig = envInitialConfig
	}

	if envDatabaseURI := os.Getenv("DATABASE_URI"); envDatabaseURI != "" {
		cfg.DatabaseDsn = envDatabaseURI
	}

	if envMigrationPath := os.Getenv("MIGRATION_PATH"); envMigrationPath != "" {
		cfg.MigrationPath = envMigrationPath
	}

	if envStand := os.Getenv("STAND"); envStand != "" {
		cfg.Stand = envStand
	}

	if envLoggerLevel := os.Getenv("LOGGER_LEVEL"); envLoggerLevel != "" {
		cfg.LoggerLevel = envLoggerLevel
	}

	if envGrpcAddress := os.Getenv("GRPC_ADDRESS"); envGrpcAddress != "" {
		cfg.GrpcAddress = envGrpcAddress
	}

	if envSslCert := os.Getenv("SSL_CERT"); envSslCert != "" {
		cfg.SslCert = envSslCert
	}

	if envSslKey := os.Getenv("SSL_KEY"); envSslKey != "" {
		cfg.SslKey = envSslKey
	}

	return nil
}

func readConfigFileServer(path string, cfg *server) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			//TODO add log
			return
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		return ErrorConfig.New("error parsing configuration file", err)
	}

	return nil
}
