package restaurants

import (
	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.RestaurantStore
}

func NewHandler(store types.RestaurantStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Get("/restaurants", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request for restaurants!")
	})
}
