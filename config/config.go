package config

import (
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	url string
}

func NewDBConfig() *DBConfig {
	url := getEnvValue("DATABASE_URL", "localhost:5432", nil)
	return &DBConfig{url: url}
}

type AppConfig struct {
	Port string
}

func NewAppConfig() *AppConfig {
	port := getEnvValue("PORT", "8080", func(value string) error {
		if portNum, err := strconv.Atoi(value); err != nil || portNum < 0 || portNum > 65535 {
			return err
		}
		return nil
	})

	return &AppConfig{Port: port}
}

func getEnvValue(envKey, defaultValue string, fn func(value string) error) string {
	value := os.Getenv(envKey)
	if value == "" {
		log.Printf("Unable load variable %s, using default value: %s", envKey, defaultValue)
		return defaultValue
	}

	if fn != nil {
		if err := fn(value); err != nil {
			log.Printf("Wrong value in envirenment variable %s, using default value: %s", envKey, defaultValue)
			return defaultValue
		}
	}

	return value
}
