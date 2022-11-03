package application_commands

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rocky.my.id/git/mygram/delivery/cli/commands"
	"rocky.my.id/git/mygram/delivery/http/api"
	"rocky.my.id/git/mygram/infrastructure/configurations/config"
	"rocky.my.id/git/mygram/infrastructure/configurations/config/keys"
	"rocky.my.id/git/mygram/infrastructure/database/common/connections"
	"syscall"
	"time"
)

var logger *zap.Logger

var ServeCmd = &cobra.Command{
	Use:   "app:serve",
	Short: "Serve the application",
	Run: func(cmd *cobra.Command, args []string) {
		ServeApp()
	},
}

func ServeApp() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	var (
		engine       = echo.New()
		db           = database_connections.Init()
		host, port   = viper.GetString(config_keys.HostAddress), viper.GetString(config_keys.HostPort)
		serveAddress = fmt.Sprintf("%s:%s", host, port)
	)

	setupEchoEngine(engine, serveAddress)
	defer syncZapLogger(logger)

	api.SetupDefault(engine, db)

	go func() {
		fmt.Println(commands.ASCIIBanner)
		if err := engine.Start(serveAddress); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server, error: " + err.Error())
		}
	}()

	<-ctx.Done()
	fmt.Println()
	fmt.Println("Shutting down gracefully, press Ctrl+C again to force")
	stop()

	{
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := engine.Shutdown(shutdownCtx); err != nil {
			log.Fatal("Server forced to shutdown: ", err)
		}
	}

	fmt.Println("Server has been shutdown.")
}

func setupEchoEngine(engine *echo.Echo, serveAddress string) {
	debug := viper.GetBool(config_keys.Debug)
	engine.HideBanner = true
	engine.Debug = debug
	engine.Server = &http.Server{
		Addr:              serveAddress,
		Handler:           engine,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	setupLogger(engine, debug)
}

func setupLogger(engine *echo.Echo, debug bool) {
	if config.IsInDevelopment() || debug {
		logger, _ = zap.NewDevelopment()
	}
	if config.IsInProduction() {
		logger, _ = zap.NewProduction()
	}
	if logger == nil {
		return
	}

	engine.Use(
		middleware.RequestLoggerWithConfig(
			middleware.RequestLoggerConfig{
				LogURI:     true,
				LogStatus:  true,
				LogMethod:  true,
				LogLatency: true,
				LogError:   true,
				LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
					logger.Info("incoming request",
						zap.String("URI", v.URI),
						zap.String("Method", v.Method),
						zap.Int("Status", v.Status),
						zap.Duration("Latency", v.Latency),
					)

					if v.Error != nil {
						logger.Error("request error",
							zap.Error(v.Error),
						)
					}

					return nil
				},
			},
		),
	)
}

func syncZapLogger(logger *zap.Logger) {
	if logger != nil {
		_ = logger.Sync()
	}
}
