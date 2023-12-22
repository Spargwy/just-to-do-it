package config

import (
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Database      Database
	Server        Server
	Logger        Logger
	Authenticator Authenticator
	Encrypter
}

type Database struct {
	Client          string `envconfig:"CLIENT_DB"`
	EnableDBLog     bool   `envconfig:"ENABLE_DB_LOG"`
	EnableLongDBLog bool   `envconfig:"ENABLE_LONG_DB_LOG"`
	MaxRetries      int    `envconfig:"MAX_RETRIES"`
}

type Server struct {
	Port string `envconfig:"PORT"`
}

type Logger struct {
	Level string `envconfig:"LOG_LEVEL"`
}

type Authenticator struct {
	JwtPath string `envconfig:"JWT_PATH"`
}

type Encrypter struct {
	Cost int `envconfig:"BCRYPT_COST"`
}

func LoadConfig(path string) *Config {
	c := filepath.Join(path, ".env")
	err := godotenv.Load(c)
	if err != nil {
		logrus.Fatalf("error load .env file, err: %v", err)
	}

	config := Config{}
	err = envconfig.Process("", &config)
	if err != nil {
		logrus.Fatalf("error process .env file, err: %v", err)
	}

	return &config
}
