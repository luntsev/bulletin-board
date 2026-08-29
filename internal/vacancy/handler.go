package vacancy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, logger *zerolog.Logger) *VacancyHandler {
	handler := &VacancyHandler{
		router:       router,
		customLogger: logger,
	}

	vacancyGroup := handler.router.Group("/vacancy")
	vacancyGroup.Post("/", handler.createVacancy)

	api := router.Group("/api")
	api.Post("/", handler.createVacancy)
	return handler
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")
	h.customLogger.Info().Msg(email)
	return c.SendString("create vacancy")
}
