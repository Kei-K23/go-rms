package restaurants

import (
	"net/http"
	"strconv"

	"github.com/Kei-K23/go-rms/backend/internal/db/middleware"
	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	rStore types.RestaurantStore
	uStore types.UserStore
}

func NewHandler(rStore types.RestaurantStore, uStore types.UserStore) *Handler {
	return &Handler{rStore: rStore, uStore: uStore}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Post("/restaurants", h.createRestaurant)
	router.Put("/restaurants/:id", h.updateRestaurant)
	router.Delete("/restaurants/:id", h.deleteRestaurant)
	router.Get("/restaurants/:id", h.getRestaurantByID)
}

func (h *Handler) createRestaurant(c *fiber.Ctx) error {
	var payload types.CreateRestaurant
	uID := c.Context().UserValue(middleware.ClaimsContextKey).(int)

	if err := utils.ParseJson(c, &payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	u, err := h.uStore.GetUserById(uID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}
	payload.AccessToken = u.AccessKey

	r, err := h.rStore.CreateRestaurant(payload)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusCreated, r)
}

func (h *Handler) updateRestaurant(c *fiber.Ctx) error {
	var payload types.UpdateRestaurant
	uID := c.Context().UserValue(middleware.ClaimsContextKey).(int)
	rID := c.Params("id")

	intRID, err := strconv.Atoi(rID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	if err := utils.ParseJson(c, &payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	u, err := h.uStore.GetUserById(uID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	r, err := h.rStore.UpdateRestaurant(payload, u.AccessKey, intRID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, r)
}

func (h *Handler) getRestaurantByID(c *fiber.Ctx) error {
	uID := c.Context().UserValue(middleware.ClaimsContextKey).(int)
	rID := c.Params("id")

	intRID, err := strconv.Atoi(rID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	u, err := h.uStore.GetUserById(uID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	r, err := h.rStore.GetRestaurantByID(intRID, u.AccessKey)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, r)
}

func (h *Handler) deleteRestaurant(c *fiber.Ctx) error {
	uID := c.Context().UserValue(middleware.ClaimsContextKey).(int)
	rID := c.Params("id")

	intRID, err := strconv.Atoi(rID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	u, err := h.uStore.GetUserById(uID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	r, err := h.rStore.DeleteRestaurant(intRID, u.AccessKey)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, r)
}
