package initialize

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	// AppEnv        string `mapstructure:"APP_ENV"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	// DBDriver      string `mapstructure:"DB_DRIVER"`
	// AppVersion    string `mapstructure:"APP_VERSION"`
	// ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from a specific .env file.
func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env_dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		// If the config file is not found, return a specific error
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return config, fmt.Errorf("config file not found: %w", err)
		}
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("unable to decode config into struct: %w", err)
	}
	return
}
