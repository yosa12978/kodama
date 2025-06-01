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
	"github.com/yosa12978/kodama/internal/server"
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

	appServer := server.New(a.config.App.Addr)

	errCh := make(chan error, 1)
	go func() {
		if err := appServer.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	slog.Info("Server is running", "addr", appServer.Addr)

	var err error
	select {
	case err = <-errCh:
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = appServer.Shutdown(timeout)
	}

	return err
}
