package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresUser        string `mapstructure:"POSTGRES_USER"`
	PostgresPassword    string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB          string `mapstructure:"POSTGRES_DB"`
	DatabaseURL         string `mapstructure:"DATABASE_URL"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	return config, err
}