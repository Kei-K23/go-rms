package orderitem

import (
	"net/http"

	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
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
	router.Put("/restaurants/:restaurantId/orderitems/:orderItemId", h.changeOrderItemStatus)
	router.Post("/restaurants/:restaurantId/orderitems", h.createOrderItem)
}

func (h *Handler) createOrderItem(c *fiber.Ctx) error {
	var payload types.CreateOrderItem
	if err := utils.ParseJson(c, &payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	oi, err := h.store.CreateOrderItem(payload)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	err = h.store.UpdateOrderStatus(oi.OrderID, "pending")
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	err = h.store.UpdateOrder(oi.OrderID, oi.Price, oi.Quantity)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusCreated, oi)
}

func (h *Handler) changeOrderItemStatus(c *fiber.Ctx) error {
	oiID, err := c.ParamsInt("orderItemId")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	res, err := h.store.ChangeOrderItemStatus(oiID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, res)
}
