package configs

import (
	"errors"

	"github.com/spf13/viper"
)

// Config App config struct
type Config struct {
	Server Server `mapstructure:"server"`
}

// Server config struct
type Server struct {
	AppVersion string `mapstructure:"app-version"`
	Port       string `mapstructure:"port"`
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
