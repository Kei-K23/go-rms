package users

import (
	"net/http"

	"github.com/Kei-K23/go-rms/backend/internal/db/middleware"
	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Get("/users", h.getUser)
}

func (h *Handler) getUser(c *fiber.Ctx) error {
	uID := c.Context().UserValue(middleware.ClaimsContextKey).(int)

	u, err := h.store.GetUserById(uID)
	if err != nil {
		return utils.WriteError(c, http.StatusUnauthorized, err)
	}

	return utils.WriteJSON(c, http.StatusOK, map[string]string{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"phone":      u.Phone,
		"address":    u.Address,
		"access_key": u.AccessKey,
		"created_at": u.CreatedAt,
	})
}
