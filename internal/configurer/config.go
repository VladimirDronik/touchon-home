package configurer

import (
	"os"
	"strconv"
	"time"
)

// Config ...
type Config struct {
	BindAddr     string        `toml:"BIND_ADDR"`
	LogLevel     string        `toml:"LOG_LEVEL"`
	DatabaseURL  string        `toml:"DATABASE_URL"`
	TokenSecret  string        `toml:"TOKEN_SECRET"`
	MaxLifetime  time.Duration `toml:"MAX_LIFETIME"`
	MaxIDLETime  time.Duration `toml:"MAX_IDLE_TIME"`
	MaxOpenConns int           `toml:"MAX_OPEN_CONNS"`
	MaxIDLEConns int           `toml:"MAX_IDLE_CONNS"`
}

// Конструктор конфигурации, присвоение дефолнтных значений
func NewConfig() *Config {

	bindAddr, _ := os.LookupEnv("BIND_ADDR")
	logLevel, _ := os.LookupEnv("LOG_LEVEL")
	databaseURL, _ := os.LookupEnv("DATABASE_URL")

	tokenSecret, _ := os.LookupEnv("TOKEN_SECRET")

	maxLifetimeString, _ := os.LookupEnv("MAX_LIFETIME")
	maxLifetime, _ := time.ParseDuration(maxLifetimeString)

	maxIDLETimeString, _ := os.LookupEnv("MAX_IDLE_TIME")
	maxIDLETime, _ := time.ParseDuration(maxIDLETimeString)

	maxOpenConnsString, _ := os.LookupEnv("MAX_OPEN_CONNS")
	maxOpenConns, _ := strconv.Atoi(maxOpenConnsString)

	maxIDLEConnsString, _ := os.LookupEnv("MAX_IDLE_CONNS")
	maxIDLEConns, _ := strconv.Atoi(maxIDLEConnsString)

	return &Config{
		BindAddr:     bindAddr,
		LogLevel:     logLevel,
		DatabaseURL:  databaseURL,
		TokenSecret:  tokenSecret,
		MaxLifetime:  maxLifetime,
		MaxIDLETime:  maxIDLETime,
		MaxOpenConns: maxOpenConns,
		MaxIDLEConns: maxIDLEConns,
	}

}
