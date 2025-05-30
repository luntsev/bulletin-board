package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"github.com/rs/zerolog"
)

type DBConfig struct {
	Url string
}

func NewDBConfig() *DBConfig {
	url := getEnvValue("DATABASE_URL", "localhost:5432", nil)
	return &DBConfig{Url: url}
}

type LogConfig struct {
	LogLevel  zerolog.Level
	LogFormat string
}

func NewLogConfig() *LogConfig {
	var logConf LogConfig

	logConf.LogFormat = getEnvValue("LOG_FORMAT", "console", func(value string) error {
		if value != "json" && value != "console" {
			return errors.New("wrong value in variable LOG_FORMAT")
		}
		return nil
	})

	logLevelStr := getEnvValue("LOG_LEVEL", "warning", func(value string) error {
		if value != "trace" && value != "debug" && value != "info" && value != "warning" && value != "error" && value != "fatal" && value != "panic" {
			return errors.New("wrong value in variable LOG_LEVEL")
		}
		return nil
	})

	switch logLevelStr {
	case "trace":
		logConf.LogLevel = zerolog.TraceLevel
	case "debug":
		logConf.LogLevel = zerolog.DebugLevel
	case "info":
		logConf.LogLevel = zerolog.InfoLevel
	case "warning":
		logConf.LogLevel = zerolog.WarnLevel
	case "error":
		logConf.LogLevel = zerolog.ErrorLevel
	case "fatal":
		logConf.LogLevel = zerolog.FatalLevel
	case "panic":
		logConf.LogLevel = zerolog.PanicLevel
	default:
		logConf.LogLevel = zerolog.NoLevel
	}

	return &logConf
}

type AppConfig struct {
	Port     string
	LogLevel zerolog.Level
}

func NewAppConfig() *AppConfig {
	port := getEnvValue("PORT", "8080", func(value string) error {
		if portNum, err := strconv.Atoi(value); err != nil || portNum < 0 || portNum > 65535 {
			return err
		}
		return nil
	})

	return &AppConfig{Port: ":" + port}
}

func getEnvValue(envKey, defaultValue string, fn func(value string) error) string {
	value := os.Getenv(envKey)
	if value == "" {
		log.Warnf("Unable load variable %s, using default value: %s", envKey, defaultValue)
		return defaultValue
	}

	if fn != nil {
		if err := fn(value); err != nil {
			log.Warnf("Wrong value in envirenment variable %s, using default value: %s", envKey, defaultValue)
			return defaultValue
		}
	}

	return value
}
