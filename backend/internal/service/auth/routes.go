package auth

import (
	"net/http"

	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	uStore types.UserStore
	aStore types.AuthStore
}

func NewHandler(uStore types.UserStore, aStore types.AuthStore) *Handler {
	return &Handler{uStore: uStore, aStore: aStore}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Post("/register", h.registerUser)
}

func (h *Handler) registerUser(c *fiber.Ctx) error {
	var payload types.RegisterUser

	if err := utils.ParseJson(c, &payload); err != nil {
		return err
	}

	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	hashedPassword, err := h.aStore.HashedPassword(payload.Password)

	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	_, err = h.uStore.CreateUser(types.RegisterUser{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		Address:  payload.Address,
		Phone:    payload.Phone,
	})

	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusCreated, map[string]string{
		"name":    payload.Name,
		"email":   payload.Email,
		"phone":   payload.Phone,
		"address": payload.Address,
	})
}
