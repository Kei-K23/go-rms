package menus

import (
	"fmt"
	"net/http"

	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.MenuStore
}

func NewHandler(store types.MenuStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Post("/restaurants/:restaurantId/menus", h.createMenu)
}

func (h *Handler) createMenu(c *fiber.Ctx) error {
	rID, err := c.ParamsInt("restaurantId")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	var payload types.CreateMenu
	if err := utils.ParseJson(c, &payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	fmt.Println(payload)

	m, err := h.store.CreateMenu(payload, rID)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	return utils.WriteJSON(c, http.StatusCreated, m)
}
