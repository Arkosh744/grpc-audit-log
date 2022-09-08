package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBURI      string `mapstructure:"DB_URI"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBDatabase string `mapstructure:"DB_DATABASE"`
	ServerPort int    `mapstructure:"SERVER_PORT"`
}

func New(folder string) (*Config, error) {
	cfg := new(Config)

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(folder)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
