package api

import (
	"database/sql"
	"log"

	"github.com/Kei-K23/go-rms/backend/internal/service/staff"
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

	// logger middleware
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	staffStore := staff.NewStore(s.db)

	staffHandler := staff.NewHandler(staffStore)

	staffHandler.RegisterRoute(v1)
	log.Fatal(app.Listen(s.addr))
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
