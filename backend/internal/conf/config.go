package conf

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigInterface interface {
	CreateClientDatabase() (interface{}, interface{}, error)
}

type Config struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_SSL      string
	POSTGRES_TIMEZONE string
	FE_URL            string
	ENV               string
	PORT              int
	AUTO_MIGRATE      bool
	SeedOnBoot        bool
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrThrow(key string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func getEnvBoolOrDefault(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	switch value {
	case "true", "TRUE", "1", "yes", "YES", "on", "ON":
		return true
	case "false", "FALSE", "0", "no", "NO", "off", "OFF":
		return false
	default:
		log.Fatalf("Invalid value for boolean environment variable %s: %s", key, value)
		return defaultValue
	}
}

func NewConfig() *Config {
	paths := []string{
		".env",
		filepath.Join("..", "..", ".env"),
	}
	loaded := false
	for _, p := range paths {
		if err := godotenv.Load(p); err == nil {
			log.Printf("Loaded environment variables from %s", p)
			loaded = true
			break
		}
	}
	if !loaded {
		log.Printf("Warning: .env file not found in %v, using OS environment", paths)
	}

	portStr := getEnvOrDefault("PORT", "9090")
	portInt := 9090
	if p, err := strconv.Atoi(portStr); err == nil {
		portInt = p
	} else {
		log.Printf("Warning: Invalid PORT value '%s', using default 9090", portStr)
	}

	return &Config{
		POSTGRES_USER:     getEnvOrThrow("POSTGRES_USER"),
		POSTGRES_PASSWORD: getEnvOrThrow("POSTGRES_PASSWORD"),
		POSTGRES_DB:       getEnvOrThrow("POSTGRES_DB"),
		POSTGRES_HOST:     getEnvOrDefault("POSTGRES_HOST", "postgres"),
		POSTGRES_PORT:     getEnvOrDefault("POSTGRES_PORT", "5432"),
		POSTGRES_SSL:      getEnvOrDefault("POSTGRES_SSL", "disable"),
		POSTGRES_TIMEZONE: getEnvOrDefault("POSTGRES_TIMEZONE", "Asia/Bangkok"),
		ENV:               getEnvOrDefault("ENV", "dev"),
		FE_URL:            getEnvOrDefault("FE_URL", "http://localhost:3000"),
		PORT:              portInt,
		AUTO_MIGRATE:      getEnvBoolOrDefault("AUTO_MIGRATE", false),
		SeedOnBoot:        getEnvBoolOrDefault("SEED_ON_BOOT", false),
	}

}
