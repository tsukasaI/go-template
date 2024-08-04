package main

import (
	"context"
	"fmt"
	"log/slog"
	"my-go-template/src/driver/web"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	engine := gin.Default()
	r := web.NewEngine(engine)
	r.SetupRouter()

	go func() {
		if err := r.Engine.Run(); err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("failed to start server: %s", err.Error()))
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// graceful shutdown
	slog.Info("graceful shutdown complete")

}

func SampleMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
