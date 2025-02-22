package handlers

import (
	"log/slog"
	"net/http"
	"time"

	headersmiddleware "github.com/DimTur/empty_service/internal/handlers/middleware/headers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/go-playground/validator/v10"
)

type RouterConfigurator interface {
	ConfigureRouter() http.Handler
}

type ChiRouterConfigurator struct {
	Logger    *slog.Logger
	validator *validator.Validate
}

func NewChiRouterConfigurator(
	logger *slog.Logger,
	validator *validator.Validate,
) *ChiRouterConfigurator {
	return &ChiRouterConfigurator{
		Logger:    logger,
		validator: validator,
	}
}

func (c *ChiRouterConfigurator) ConfigureRouter() http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(httprate.LimitByIP(100, 1*time.Minute))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-User-ID"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(headersmiddleware.SecurityHeadersMiddleware)

	return router
}
