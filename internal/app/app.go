package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yosa12978/kodama/internal/config"
	"github.com/yosa12978/kodama/internal/logger"
	"github.com/yosa12978/kodama/internal/templates"
)

type App struct {
	config config.Config
}

func NewFromConfig(conf config.Config) *App {
	return &App{config: conf}
}

func (a *App) Run() error {
	ctx, cancel := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	slog.SetDefault(
		logger.New(&logger.LoggerOptions{
			Sink:  os.Stdout,
			Level: a.config.App.LogLevel,
		}),
	)

	server := http.Server{
		Addr: a.config.App.Addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			templates.IndexTemplate.Execute(w, nil)
		}),
	}

	errCh := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	slog.Info("Server is running", "addr", server.Addr)

	var err error
	select {
	case err = <-errCh:
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = server.Shutdown(timeout)
	}

	return err
}
