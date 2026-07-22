package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Load reads configuration from environment variables
// and returns the application configuration.
func Load() (*Config, error) {

	// Read values from .env file
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	// Environment variables override .env values
	viper.AutomaticEnv()

	// Load .env if available
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Warning: .env file not found. Using environment variables.")
	}

	cfg := &Config{
		App: AppConfig{
			Name:         viper.GetString("APP_NAME"),
			Env:          viper.GetString("APP_ENV"),
			Port:         viper.GetString("PORT"),
			LogLevel:     viper.GetString("LOG_LEVEL"),
			ReadTimeout:  viper.GetInt("READ_TIMEOUT"),
			WriteTimeout: viper.GetInt("WRITE_TIMEOUT"),
		},

		DB: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
			SSLMode:  viper.GetString("DB_SSLMODE"),
			TimeZone: viper.GetString("DB_TIMEZONE"),
		},
	}

	// Basic App Validation
	if cfg.App.Name == "" {
		return nil, fmt.Errorf("APP_NAME is required")
	}

	if cfg.App.Env == "" {
		return nil, fmt.Errorf("APP_ENV is required")
	}

	if cfg.App.Port == "" {
		return nil, fmt.Errorf("PORT is required")
	}

	if cfg.App.LogLevel == "" {
		cfg.App.LogLevel = "info"
	}

	if cfg.App.ReadTimeout <= 0 {
		cfg.App.ReadTimeout = 10
	}

	if cfg.App.WriteTimeout <= 0 {
		cfg.App.WriteTimeout = 10
	}

	// Basic DB Validation
	if cfg.DB.Host == "" {
		return nil, fmt.Errorf("DB_HOST is required")
	}

	if cfg.DB.Port == "" {
		cfg.DB.Port = "5432"
	}

	if cfg.DB.SSLMode == "" {
		cfg.DB.SSLMode = "disable"
	}

	if cfg.DB.TimeZone == "" {
		cfg.DB.TimeZone = "UTC"
	}
	return cfg, nil
}
