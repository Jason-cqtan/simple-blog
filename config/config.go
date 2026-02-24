package config

import (
	"log"
	"os"
)

type Config struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	ServerPort string
	SecureCookie bool
}

const defaultJWTSecret = "secret-key-change-in-production"

func LoadConfig() *Config {
	jwtSecret := getEnv("JWT_SECRET", defaultJWTSecret)
	if jwtSecret == defaultJWTSecret {
		log.Println("WARNING: Using default JWT secret. Set JWT_SECRET environment variable for production.")
	}

	return &Config{
		DBDriver:     getEnv("DB_DRIVER", "mysql"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "3306"),
		DBUser:       getEnv("DB_USER", "root"),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_NAME", "simple_blog"),
		JWTSecret:    jwtSecret,
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		SecureCookie: getEnv("SECURE_COOKIE", "true") != "false",
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
