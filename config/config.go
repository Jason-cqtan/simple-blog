package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Config struct {
	DBDriver     string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecret    string
	ServerPort   string
	SecureCookie bool
}

const defaultJWTSecret = "secret-key-change-in-production"

// loadDotEnv reads a .env file and sets environment variables.
// Existing environment variables are not overwritten.
func loadDotEnv(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		// .env file is optional; silently skip if not found
		return
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		// Remove surrounding quotes if present
		if len(value) >= 2 &&
			((value[0] == '"' && value[len(value)-1] == '"') ||
				(value[0] == '\'' && value[len(value)-1] == '\'')) {
			value = value[1 : len(value)-1]
		}
		// Do not overwrite existing environment variables
		if os.Getenv(key) == "" {
			_ = os.Setenv(key, value)
		}
	}
}

func LoadConfig() *Config {
	loadDotEnv(".env")

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
