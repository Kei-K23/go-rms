package api

import (
	"database/sql"
	"log"

	"github.com/Kei-K23/go-rms/backend/internal/db/middleware"
	"github.com/Kei-K23/go-rms/backend/internal/service/auth"
	"github.com/Kei-K23/go-rms/backend/internal/service/staff"
	"github.com/Kei-K23/go-rms/backend/internal/service/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func (s *APIServer) Run() {
	app := fiber.New()
	// global middleware
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	protectedRoute := v1.Group("")
	// protedted route with auth middleware
	protectedRoute.Use(middleware.AuthMiddleware)

	// stores service
	staffStore := staff.NewStore(s.db)
	userStore := users.NewStore(s.db)
	authStore := auth.NewStore(s.db)

	// handlers
	staffHandler := staff.NewHandler(staffStore)
	authHandler := auth.NewHandler(userStore, authStore)
	userHandler := users.NewHandler(userStore)
	// register routes
	staffHandler.RegisterRoute(v1)
	authHandler.RegisterRoute(v1)
	userHandler.RegisterRoute(protectedRoute)

	// server
	log.Fatal(app.Listen(s.addr))
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
