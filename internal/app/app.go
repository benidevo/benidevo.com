package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/benidevo/website/internal/config"
	"github.com/benidevo/website/internal/router"
	"github.com/gin-gonic/gin"
)

// App represents the main application structure, holding configuration settings,
// the Gin router, HTTP server, and a channel for OS signals to manage graceful shutdown.
type App struct {
	cfg    *config.Config
	router *gin.Engine
	server *http.Server
	done   chan os.Signal
}

func New(cfg *config.Config) *App {
	app := &App{
		cfg:  cfg,
		done: make(chan os.Signal, 1),
	}

	return app
}

func (a *App) Setup() error {
	log.Info().Msgf("Starting server on port %s", a.cfg.Settings.Port)

	router, err := router.SetupRouter(a.cfg)
	if err != nil {
		return err
	}
	a.router = router

	log.Info().Msg("Routes have been set up successfully")

	return nil
}

func (a *App) Run() error {
	if err := a.Setup(); err != nil {
		log.Error().Err(err).Msg("Failed to set up application")
		return err
	}

	a.server = &http.Server{
		Addr:    ":" + a.cfg.Settings.Port,
		Handler: a.router,
	}

	signal.Notify(a.done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}

	}()

	return nil
}

func (a *App) WaitForShutdown() {
	sig := <-a.done
	log.Info().Msgf("Received signal: %s. Shutting down server...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to gracefully shutdown server")
	} else {
		log.Info().Msg("Server shutdown gracefully")
	}

}
