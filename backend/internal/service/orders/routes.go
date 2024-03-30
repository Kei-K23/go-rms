package orders

import (
	"net/http"

	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.OrderStore
}

type createOrderPayload struct {
	Order      types.CreateOrder       `json:"order" validate:"required"`
	OrderItems []types.CreateOrderItem `json:"order_items" validate:"required"`
}

func NewHandler(s types.OrderStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Post("/restaurants/:restaurantId/orders", h.createOrder)
}

//
// func (h *Handler) getOrdersByRestaurantId(c *fiber.Ctx) error {
// }

func (h *Handler) createOrder(c *fiber.Ctx) error {
	var payload createOrderPayload
	if err := utils.ParseJson(c, &payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	order, err := h.store.CreateOrder(payload.Order, payload.OrderItems)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}
	return utils.WriteJSON(c, http.StatusCreated, order)
}
