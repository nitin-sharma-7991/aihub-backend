package config

import "time"

// AppConfig contains application specific configuration.
type AppConfig struct {
	Name         string
	Env          string
	Port         string
	LogLevel     string
	ReadTimeout  int
	WriteTimeout int
}

// Database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

type JWTConfig struct {
	Secret    string
	ExpiresIn time.Duration
}

// Root Config
// Config is the root configuration object.
// Future configurations like Database, Redis, JWT, etc.
// will be added here.
type Config struct {
	App AppConfig
	DB  DatabaseConfig
	JWT JWTConfig
}
