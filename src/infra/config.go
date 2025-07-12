package infra

import (
	"os"
	"strconv"
)

type SQLConfig struct {
	Host           string
	Port           int
	User           string
	Password       string
	DBName         string
	SSLMode        string
	DBInstanceType string
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

var PostgresCredentialsConfig = SQLConfig{
	Host:           getEnvOrDefault("DB_HOST", "postgres"),
	Port:           getEnvOrDefaultInt("DB_PORT", 5432),
	User:           getEnvOrDefault("DB_USER", "user"),
	Password:       getEnvOrDefault("DB_PASSWORD", "password"),
	DBName:         getEnvOrDefault("DB_NAME", "db"),
	SSLMode:        getEnvOrDefault("DB_SSLMODE", "disable"),
	DBInstanceType: "postgres",
}
