package main

import (
	"bulletin-board/config"
	"bulletin-board/internal/home"
	"bulletin-board/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Warnf("Unable load envirenmenl variables from .env file: %v", err)
	}
}

func main() {
	logConf := config.NewLogConfig()
	customLogger := logger.NewLogger(logConf)

	app := fiber.New()
	appConf := config.NewAppConfig()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))

	app.Use(recover.New())
	home.NewHomeHandler(app, customLogger)

	app.Listen(appConf.Port)
}
