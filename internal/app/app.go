package app

import (
	"log/slog"
	"time"

	"github.com/DimTur/empty_service/internal/app/httpapp"
	"github.com/DimTur/empty_service/internal/handlers"
	"github.com/go-playground/validator/v10"
)

type App struct {
	HTTPServer *httpapp.APIServer
}

func NewApp(
	httpAddr string,
	readTimeout time.Duration,
	writeTimeout time.Duration,
	iddleTimeout time.Duration,
	logger *slog.Logger,
	validator *validator.Validate,
) (*App, error) {
	routerConfigurator := handlers.NewChiRouterConfigurator(
		logger,
		validator,
	)
	router := routerConfigurator.ConfigureRouter()

	httpServer, err := httpapp.NewHTTPServer(
		httpAddr,
		router,
		readTimeout,
		writeTimeout,
		iddleTimeout,
		logger,
		validator,
	)
	if err != nil {
		logger.Error("failed to create server", slog.Any("err", err))
		return nil, err
	}

	return &App{
		HTTPServer: httpServer,
	}, nil
}
