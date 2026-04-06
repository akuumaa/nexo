package config

import "os"

type Config struct {
	AppEnv        string
	AppHost       string
	AppPort       string
	DBHost        string
	DBPort        string
	DBName        string
	DBUser        string
	DBPassword    string
	DBSSLMode     string
	SessionSecret string
}

func Load() Config {
	return Config{
		AppEnv:        getEnv("APP_ENV", "development"),
		AppHost:       getEnv("APP_HOST", "0.0.0.0"),
		AppPort:       getEnv("APP_PORT", "8080"),
		DBHost:        getEnv("DB_HOST", "db"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBName:        getEnv("DB_NAME", "nexo"),
		DBUser:        getEnv("DB_USER", "nexo"),
		DBPassword:    getEnv("DB_PASSWORD", "nexo"),
		DBSSLMode:     getEnv("DB_SSL_MODE", "disable"),
		SessionSecret: getEnv("SESSION_SECRET", "change-me"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
