package config

// AppConfig contains application specific configuration.
type AppConfig struct {
	Name         string
	Env          string
	Port         string
	LogLevel     string
	ReadTimeout  int
	WriteTimeout int
}

// Config is the root configuration object.
// Future configurations like Database, Redis, JWT, etc.
// will be added here.
type Config struct {
	App AppConfig
}
