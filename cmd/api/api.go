package api

import (
	"database/sql"
	"log"

	"github.com/Kei-K23/go-rms/backend/internal/config"
	"github.com/Kei-K23/go-rms/backend/internal/db/middleware"
	"github.com/Kei-K23/go-rms/backend/internal/service/auth"
	"github.com/Kei-K23/go-rms/backend/internal/service/category"
	"github.com/Kei-K23/go-rms/backend/internal/service/menus"
	orderitem "github.com/Kei-K23/go-rms/backend/internal/service/orderItem"
	"github.com/Kei-K23/go-rms/backend/internal/service/orders"
	"github.com/Kei-K23/go-rms/backend/internal/service/payment"
	"github.com/Kei-K23/go-rms/backend/internal/service/restaurantTables"
	"github.com/Kei-K23/go-rms/backend/internal/service/restaurants"
	"github.com/Kei-K23/go-rms/backend/internal/service/staff"
	"github.com/Kei-K23/go-rms/backend/internal/service/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/stripe/stripe-go/v76"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func (s *APIServer) Run() {
	stripe.Key = config.Env.STRIPE_API_KEY

	app := fiber.New()
	// global middleware
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	// stores service
	staffStore := staff.NewStore(s.db)
	userStore := users.NewStore(s.db)
	authStore := auth.NewStore(s.db)
	restStore := restaurants.NewStore(s.db)
	restTStore := restaurantTables.NewStore(s.db)
	categoryStore := category.NewStore(s.db)
	menuStore := menus.NewStore(s.db)
	orderItemStore := orderitem.NewStore(s.db)
	orderStore := orders.NewStore(s.db, orderItemStore)
	paymentStore := payment.NewStore(s.db)

	// handlers
	staffHandler := staff.NewHandler(staffStore)
	authHandler := auth.NewHandler(userStore, authStore)
	userHandler := users.NewHandler(userStore)
	restHandler := restaurants.NewHandler(restStore, userStore)
	restTHandler := restaurantTables.NewHandler(restTStore)
	categoryHandler := category.NewHandler(categoryStore)
	menuHandler := menus.NewHandler(menuStore)
	orderHandler := orders.NewHandler(orderStore)
	orderItemHandler := orderitem.NewHandler(orderItemStore)
	paymentHandler := payment.NewHandler(orderItemStore, paymentStore)
	// register routes
	staffHandler.RegisterRoute(v1)
	authHandler.RegisterRoute(v1)
	paymentHandler.RegisterRoute(v1)

	// protedted route with auth middleware
	protectedRoute := v1.Group("")
	protectedRoute.Use(middleware.AuthMiddleware)

	userHandler.RegisterRoute(protectedRoute)
	restHandler.RegisterRoute(protectedRoute)
	restTHandler.RegisterRoute(protectedRoute)
	categoryHandler.RegisterRoute(protectedRoute)
	menuHandler.RegisterRoute(protectedRoute)
	orderHandler.RegisterRoute(protectedRoute)
	orderItemHandler.RegisterRoute(protectedRoute)

	// server
	log.Fatal(app.Listen(s.addr))
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
