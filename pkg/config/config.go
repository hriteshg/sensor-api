package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"sensor-api/pkg/db"
)

type Config struct {
	Environment   string `envconfig:"ENVIRONMENT"`
	PORT          string `envconfig:"PORT" default:"3333"`
	MigrationPath string `envconfig:"MIGRATION_PATH" default:"file:///app/migration"`
	DBHost        string `envconfig:"DB_HOST"`
	DBDriver      string `envconfig:"DB_DRIVER"`
	DBUser        string `envconfig:"DB_USER"`
	DBPassword    string `envconfig:"DB_PASSWORD"`
	DBName        string `envconfig:"DB_NAME"`
	DBPort        string `envconfig:"DB_PORT"`
	RedisUrl      string `envconfig:"REDIS_URL"`
}

func Init() Config {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Warn("Failed to load environment variables from .env file")
	}

	var config Config
	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	return config
}

func (c Config) DBConfig() db.DatabaseConfig {
	return db.DatabaseConfig{
		Name:     c.DBName,
		Host:     c.DBHost,
		Port:     c.DBPort,
		Username: c.DBUser,
		Password: c.DBPassword,
	}
}
