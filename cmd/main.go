package main

import (
	"bulletin-board/config"
	"bulletin-board/internal/home"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Unable load envirenmenl variables from .env file: %v", err)
	}
}

func main() {
	app := fiber.New()
	appConf := config.NewAppConfig()
	app.Use(recover.New())
	home.NewHomeHandler(app)

	app.Listen(appConf.Port)
}
