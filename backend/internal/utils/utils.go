package utils

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func ParseJson(c *fiber.Ctx, payload any) error {
	if len(c.Body()) <= 0 {
		return fmt.Errorf("body payload is missing")
	}
	if err := c.BodyParser(&payload); err != nil {
		return fmt.Errorf("error when parsing payload")
	}

	return nil
}

func WriteJSON(c *fiber.Ctx, status int, payload any) error {
	c.Status(status)
	c.Response().Header.Set("Content-Type", "application/json")

	return c.JSON(payload)
}

func ValidatePayload(payload any) error {
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}

func WriteError(c *fiber.Ctx, status int, err error) error {
	return WriteJSON(c, status, map[string]string{"error": err.Error()})
}
