package home

import (
	"github.com/gofiber/fiber/v3"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHomeHandler(router fiber.Router) *HomeHandler {
	handler := &HomeHandler{router: router}

	api := router.Group("/api")

	api.Get("/", handler.home)

	return handler
}

func (h *HomeHandler) home(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
