package category

import (
	"fmt"
	"net/http"

	"github.com/Kei-K23/go-rms/backend/internal/types"
	"github.com/Kei-K23/go-rms/backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.CategoryStore
}

func NewHandler(store types.CategoryStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute(router fiber.Router) {
	router.Get("/categories", h.getAllCategory)
	router.Get("/categories/:id", h.getCategoryByID)
	router.Put("/categories/:id", h.updateCategory)
	router.Delete("/categories/:id", h.deleteCategory)
	router.Post("/categories", h.createCategory)
}

func (h *Handler) getAllCategory(c *fiber.Ctx) error {
	categories, err := h.store.GetCategories()
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, categories)
}

func (h *Handler) getCategoryByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	ct, err := h.store.GetCategoryByID(id)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError,
			fmt.Errorf("can't get category with id %d", id))
	}

	return utils.WriteJSON(c, http.StatusOK, ct)
}

func (h *Handler) deleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	res, err := h.store.DeleteCategory(id)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, res)
}

func (h *Handler) updateCategory(c *fiber.Ctx) error {
	var payload types.UpdateCategory
	if err := utils.ParseJson(c, &payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	ct, err := h.store.UpdateCategory(payload, id)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, ct)
}

func (h *Handler) createCategory(c *fiber.Ctx) error {
	var payload types.CreateCategory
	if err := utils.ParseJson(c, &payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	if err := utils.ValidatePayload(payload); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	ct, err := h.store.CreateCategory(payload)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusCreated, ct)
}
