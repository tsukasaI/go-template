package web

import (
	"fmt"
	"my-go-template/src/core/config"
	"my-go-template/src/di"

	"github.com/labstack/echo/v4"
)

type EchoEngine struct {
	e *echo.Echo
}

func NewEchoEngine(e *echo.Echo) *EchoEngine {
	ee := &EchoEngine{e}
	ee.setupRouter()
	return ee
}

func (ee *EchoEngine) setupRouter() {
	v1Group := ee.e.Group("/v1")

	sampleController := di.GetSampleEchoController()

	v1Group.GET("/samples", func(ctx echo.Context) error {
		return ctx.JSON(sampleController.ExecuteGetSample(ctx))
	})

}

func (ee *EchoEngine) Start() error {
	return ee.e.Start(fmt.Sprintf(":%d", config.Port()))
}
