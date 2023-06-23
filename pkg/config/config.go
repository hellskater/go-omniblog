package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold the configuration values
type Config struct {
	DBUser          string
	DBPassword      string
	DBName          string
	ServerPort      string
	RecaptchaSecret string
}

// NewConfig creates new config instance
func NewConfig() *Config {
	// loads environment variables from .env file
	err := godotenv.Load(
		"../../.env",
	)
	if err != nil {
		return nil
	}

	return &Config{
		DBUser:          getEnv("DB_USER", ""),
		DBPassword:      getEnv("DB_PASSWORD", ""),
		DBName:          getEnv("DB_NAME", ""),
		ServerPort:      getEnv("SERVER_PORT", "3000"),
		RecaptchaSecret: getEnv("RECAPTCHA_SECRET", ""),
	}
}

// helper function to read environment variables
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
