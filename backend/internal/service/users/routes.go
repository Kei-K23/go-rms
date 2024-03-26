package users

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Get("/users", h.getUser)
}

func (h *Handler) getUser(c *fiber.Ctx) error {
	return nil
}
