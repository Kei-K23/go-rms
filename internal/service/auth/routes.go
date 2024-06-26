package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kei-K23/go-rms/backend/internal/config"
	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	router.Post("/login", h.loginUser)
}

func (h *Handler) registerUser(c *fiber.Ctx) error {
	var payload types.RegisterUser
	payload.AccessKey = uuid.New().String()

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
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  hashedPassword,
		Address:   payload.Address,
		Phone:     payload.Phone,
		AccessKey: payload.AccessKey,
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

func (h *Handler) loginUser(c *fiber.Ctx) error {
	var payload types.LoginUser

	if err := utils.ParseJson(c, &payload); err != nil {
		return err
	}

	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	u, err := h.uStore.GetUserByEmail(payload)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("cannot find the user"))
	}

	err = h.aStore.VerifyPassword(u.Password, payload.Password)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	uID, err := strconv.Atoi(u.ID)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	accessKey, err := h.aStore.CreateJWT([]byte(config.Env.SECRET_KEY), uID)
	if err != nil {
		return err
	}

	return utils.WriteJSON(c, http.StatusOK, map[string]string{
		"access_key": accessKey,
	})
}
