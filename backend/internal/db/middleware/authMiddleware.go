package middleware

import (
	"net/http"
	"strings"

	"github.com/Kei-K23/go-rms/backend/internal/config"
	"github.com/Kei-K23/go-rms/backend/internal/service/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const ClaimsContextKey ContextKey = "claims"

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	// Check if Authorization header is present
	if authHeader == "" {
		return &fiber.Error{
			Code:    http.StatusUnauthorized,
			Message: "authorization header is missing",
		}
	}
	// Extract the token from the Authorization header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &auth.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.SECRET_KEY), nil
	})
	// Check for token parsing errors
	if err != nil {
		return &fiber.Error{
			Code:    http.StatusUnauthorized,
			Message: "authorization header is missing",
		}
	}

	// Validate the token
	if !token.Valid {
		return &fiber.Error{
			Code:    http.StatusUnauthorized,
			Message: "authorization header is missing",
		}
	}

	// Extract claims from the token
	claims, ok := token.Claims.(*auth.JWTClaim)

	if !ok {
		return &fiber.Error{
			Code:    http.StatusUnauthorized,
			Message: "authorization header is missing",
		}
	}
	c.Context().SetUserValue(ClaimsContextKey, claims.UserID)
	return c.Next()
}
