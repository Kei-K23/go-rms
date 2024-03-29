package restaurantTables

import (
	"net/http"
	"strconv"

	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	rTStore types.RestaurantTablesStore
}

func NewHandler(rTStore types.RestaurantTablesStore) *Handler {
	return &Handler{rTStore: rTStore}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Get("/restaurants/:id/tables", h.getTablesByRestaurant)
	router.Get("/restaurants/:id/tables/:tid", h.getTablesByRestaurantID)
	router.Put("/restaurants/:id/tables/:tid", h.updateRestaurantTable)
	router.Post("/restaurants/:id/tables", h.createRestaurantTable)
}

func (h *Handler) getTablesByRestaurant(c *fiber.Ctx) error {
	rID, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	rTables, err := h.rTStore.GetRestaurantTables(rID)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	return utils.WriteJSON(c, http.StatusCreated, rTables)
}

func (h *Handler) getTablesByRestaurantID(c *fiber.Ctx) error {
	rID, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	rTID, err := c.ParamsInt("tid")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	rTables, err := h.rTStore.GetRestaurantTableByID(rTID, rID)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	return utils.WriteJSON(c, http.StatusCreated, rTables)
}

func (h *Handler) createRestaurantTable(c *fiber.Ctx) error {
	var payload types.CreateRestaurantTable
	rID := c.Params("id")
	if err := utils.ParseJson(c, &payload); err != nil {
		return err
	}

	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	intRID, err := strconv.Atoi(rID)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	payload.RestaurantID = intRID
	rT, err := h.rTStore.CreateRestaurantTable(payload)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	return utils.WriteJSON(c, http.StatusCreated, rT)
}

func (h *Handler) updateRestaurantTable(c *fiber.Ctx) error {
	var payload types.UpdateRestaurantTable
	if err := utils.ParseJson(c, &payload); err != nil {
		return err
	}

	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	rID, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	rTID, err := c.ParamsInt("tid")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	rTables, err := h.rTStore.UpdateRestaurantTable(payload, rID, rTID)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	return utils.WriteJSON(c, http.StatusCreated, rTables)
}
