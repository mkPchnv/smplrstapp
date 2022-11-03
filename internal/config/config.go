package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	appEnv           = "APP_ENV"
	postgresHost     = "POSTGRES_HOST"
	postgresPort     = "POSTGRES_PORT"
	postgresUserName = "POSTGRES_USERNAME"
	postgresPassword = "POSTGRES_PASSWORD"
	postgresDatabase = "POSTGRES_DATABASE"
	host             = "HTTP_HOST"
	port             = "HTTP_PORT"
)

type (
	Config struct {
		AppEnv   string
		Postgres PostgresConfig
		HTTP     HTTPConfig
	}

	PostgresConfig struct {
		UserName string
		Password string
		Database string
		Host     string
		Port     string
	}

	HTTPConfig struct {
		Host string
		Port string
	}
)

func Init(pathEnvFile string) (*Config, error) {
	err := godotenv.Load(pathEnvFile)
	if err != nil {
		return nil, errors.New("error loading .env file")
	}

	config := Config{}
	setFromEnv(&config)

	return &config, nil
}

func setFromEnv(cfg *Config) {
	cfg.AppEnv = os.Getenv(appEnv)

	cfg.Postgres.UserName = os.Getenv(postgresUserName)
	cfg.Postgres.Password = os.Getenv(postgresPassword)
	cfg.Postgres.Host = os.Getenv(postgresHost)
	cfg.Postgres.Port = os.Getenv(postgresPort)
	cfg.Postgres.Database = os.Getenv(postgresDatabase)

	cfg.HTTP.Host = os.Getenv(host)
	cfg.HTTP.Port = os.Getenv(port)
}

func (pc *PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pc.UserName, pc.Password, pc.Host, pc.Port, pc.Database)
}
