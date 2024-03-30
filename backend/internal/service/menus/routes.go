package menus

import (
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
	router.Get("/restaurants/:restaurantId/menus", h.getMenusByRestaurantID)
	router.Post("/restaurants/:restaurantId/menus", h.createMenu)
	router.Get("/restaurants/:restaurantId/menus/:menuId", h.getMenuByID)
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

	m, err := h.store.CreateMenu(payload, rID)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	return utils.WriteJSON(c, http.StatusCreated, m)
}

func (h *Handler) getMenuByID(c *fiber.Ctx) error {
	restaurantID, err := c.ParamsInt("restaurantId")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	menuID, err := c.ParamsInt("menuId")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	m, err := h.store.GetMenuByID(menuID, restaurantID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}
	return utils.WriteJSON(c, http.StatusOK, m)
}

func (h *Handler) getMenusByRestaurantID(c *fiber.Ctx) error {
	restaurantID, err := c.ParamsInt("restaurantId")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	menus, err := h.store.GetAllMenuByRestaurantID(restaurantID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, menus)
}
