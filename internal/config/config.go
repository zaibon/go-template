package config

import (
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		HTTP struct {
			Port int `mapstructure:"PORT"`
		} `mapstructure:"HTTP"`
		GRPC struct {
			Port int `mapstructure:"PORT"`
		} `mapstructure:"GRPC"`
	} `mapstructure:"SERVER"`
	Log struct {
		Level string `mapstructure:"LEVEL"`
	} `mapstructure:"LOG"`
	Database struct {
		ConnectionString string `mapstructure:"CONNECTION_STRING"`
	} `mapstructure:"DATABASE"`
	SomeValue string `mapstructure:"SOME_VALUE"` // Example of a non-nested config
	// Add more configuration sections as needed
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")  // Name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()
	if err != nil {
		slog.Warn("Failed to read config file", "error", err)
	}

	err = viper.Unmarshal(&config)
	return
}
