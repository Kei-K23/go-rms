package orderitem

import (
	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.OrderItemStore
}

func NewHandler(store types.OrderItemStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
}
