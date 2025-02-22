package serve

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/DimTur/empty_service/internal/app"
	"github.com/DimTur/empty_service/internal/config"
	"github.com/DimTur/empty_service/internal/lib/validator"
	"github.com/spf13/cobra"
)

// NewServeCmd creates new serve command
func NewServeCmd() *cobra.Command {
	var configPath string

	c := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Start Empty server",
		RunE: func(cmd *cobra.Command, args []string) error {
			// logger
			log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

			ctx, cancel := signal.NotifyContext(cmd.Context(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
			defer cancel()

			// parse config
			cfg, err := config.Parse(configPath)
			if err != nil {
				return err
			}

			// validator
			validator := validator.InitValidator()

			// app
			app, err := app.NewApp(
				cfg.HTTPServer.Address,
				cfg.HTTPServer.Timeout,
				cfg.HTTPServer.Timeout,
				cfg.HTTPServer.IddleTimeout,
				log,
				validator,
			)
			if err != nil {
				return err
			}

			// run server
			httCloser, err := app.HTTPServer.Run()
			if err != nil {
				return err
			}

			log.Info("server listening:", slog.Any("port", cfg.HTTPServer.Address))
			<-ctx.Done()

			// close server
			httCloser()

			return nil
		},
	}

	// flags
	c.Flags().StringVar(&configPath, "config", "", "path to config")
	return c
}
