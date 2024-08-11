package main

import (
	"context"
	"fmt"
	"log/slog"
	"my-go-template/src/driver/db/migrate"
	"my-go-template/src/driver/web"
	"my-go-template/src/pkg/echo"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	migrate.Execute()

	echoEngine := echo.New()
	e := web.NewEchoEngine(echoEngine)

	go func() {
		if err := e.Start(); err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("failed to start server: %s", err.Error()))
			os.Exit(1)
		}
	}()

	slog.Info("server started")

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// graceful shutdown
	if err := echoEngine.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("graceful shutdown complete")

}

func SampleMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
