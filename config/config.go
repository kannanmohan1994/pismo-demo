package config

import (
	"sync"

	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Environment                      string `mapstructure:"ENVIRONMENT" validate:"required"`
	ServerPort                       string `mapstructure:"SERVER_PORT" validate:"required"`
	PostgresHost                     string `mapstructure:"POSTGRES_HOST" validate:"required"`
	PostgresDB                       string `mapstructure:"POSTGRES_DB" validate:"required"`
	PostgresSchema                   string `mapstructure:"POSTGRES_SCHEMA" validate:"required"`
	PostgresUser                     string `mapstructure:"POSTGRES_USER" validate:"required"`
	PostgresPassword                 string `mapstructure:"POSTGRES_PASSWORD" validate:"required"`
	PostgresPort                     string `mapstructure:"POSTGRES_PORT" validate:"required"`
	JWTSecretKey                     string `mapstructure:"JWT_SECRET_KEY" validate:"required"`
	AccessTokenExpiryDurationSeconds uint   `mapstructure:"ACCESS_TOKEN_EXPIRY_DURATION_SECONDS" validate:"required,gt=0"`
}

var (
	config *Config
	once   sync.Once
)

func Setup() {
	once.Do(func() {
		viper.AutomaticEnv()
		viper.SetConfigFile(".env")
		config = new(Config)
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("error reading config - %s", err)

		}
		if err := viper.Unmarshal(config); err != nil {
			log.Printf("unable to decode config - %v", err)

		}

	})
}

func GetConfig() *Config {
	return config
}
