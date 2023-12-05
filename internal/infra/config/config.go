package config

import (
	"errors"

	"github.com/spf13/viper"
)

// Config App config struct
type Config struct {
	Server   Server   `mapstructure:"server"`
	Postgres Postgres `mapstructure:"postgres"`
}

// Server config struct
type Server struct {
	AppVersion string `mapstructure:"app-version"`
	Port       string `mapstructure:"port"`
	LogLevel   string `mapstructure:"log-level"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db-name"`
}

// Load config file from given path
func Load() (config *Config, err error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	if err = v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
