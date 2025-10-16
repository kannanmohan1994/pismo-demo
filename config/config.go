package config

import (
	"sync"

	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Environment                      string `mapstructure:"ENVIRONMENT"`
	ServerPort                       string `mapstructure:"SERVER_PORT"`
	PostgresHost                     string `mapstructure:"POSTGRES_HOST"`
	PostgresDB                       string `mapstructure:"POSTGRES_DB"`
	PostgresSchema                   string `mapstructure:"POSTGRES_SCHEMA"`
	PostgresUser                     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword                 string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresPort                     string `mapstructure:"POSTGRES_PORT"`
	JWKFilePath                      string `mapstructure:"JWK_FILE_PATH"`
	JWKKid                           string `mapstructure:"JWK_KID"`
	AccessTokenExpiryDurationSeconds uint   `mapstructure:"ACCESS_TOKEN_EXPIRY_DURATION_SECONDS"`
	KafkaHost                        string `mapstructure:"KAFKA_HOST"`
	KafkaPort                        string `mapstructure:"KAFKA_PORT"`
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
