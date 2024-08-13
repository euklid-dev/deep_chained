package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	OPEN_AI_KEY      string
	DB_Host          string
	DB_Port          string
	DB_User          string
	DB_Password      string
	DB_SSLMode       string
	MASTER_DB_URL    string
	APP_SERVICE_PORT string
}

var GlobalAppConfig *AppConfig

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if validateENV() {
		GlobalAppConfig = &AppConfig{
			OPEN_AI_KEY:      extractEnv("OPENAI_API_KEY"),
			DB_Host:          extractEnv("DB_HOST"),
			DB_Port:          extractEnv("DB_PORT"),
			DB_User:          extractEnv("DB_USER"),
			DB_Password:      extractEnv("DB_PASSWORD"),
			DB_SSLMode:       extractEnv("DB_SSLMODE"),
			MASTER_DB_URL:    extractEnv("MASTER_DB_URL"),
			APP_SERVICE_PORT: extractEnv("APP_SERVICE_PORT"),
		}
	} else {
		fmt.Println("Failed to load configuration: Missing required environment variables.")
		os.Exit(1)
	}
}

func validateENV() bool {
	envs := []string{
		"OPENAI_API_KEY",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_SSLMODE",
		"MASTER_DB_URL",
		"APP_SERVICE_PORT",
	}

	for _, env := range envs {
		if os.Getenv(env) == "" {
			return false
		}
	}

	return true
}

func extractEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Failed to extract %s from environment variables.\n", key)
	}
	return value
}

// Add this method to generate a DSN for database connections
func (c *AppConfig) DSN(dbName string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DB_Host, c.DB_Port, c.DB_User, c.DB_Password, dbName, c.DB_SSLMode)
}
