package restaurants

import (
	"net/http"

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
	router.Get("/restaurants", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request for restaurants!")
	})

	router.Post("/restaurants", h.createRestaurant)
}

func (h *Handler) createRestaurant(c *fiber.Ctx) error {
	var payload types.CreateRestaurant
	uID := c.Context().UserValue(middleware.ClaimsContextKey).(int)

	if err := utils.ParseJson(c, &payload); err != nil {
		return err
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
