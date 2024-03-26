package users

import (
	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Get("/users", h.getUser)
}

func (h *Handler) getUser(c *fiber.Ctx) error {
	return nil
}
