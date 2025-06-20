package home

import (
	"bulletin-board/pkg/tadapter"
	"bulletin-board/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHomeHandler(router fiber.Router, logger *zerolog.Logger) *HomeHandler {
	handler := &HomeHandler{
		router:       router,
		customLogger: logger,
	}

	api := router.Group("/api")
	api.Get("/", handler.home)
	return handler
}

func (h *HomeHandler) home(c *fiber.Ctx) error {

	component := views.Hello("Коля")
	return tadapter.Render(c, component)

	/*h.customLogger.Warn().
	Bool("isAdmin", true).
	Str("email", "n.luntsev@yandex.ru").
	Int("IntVal", 100500).
	Msg("Log message")*/
}
