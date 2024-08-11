package controller

import (
	"my-go-template/src/adapter/controller/response"
	"my-go-template/src/usecase"

	"github.com/labstack/echo/v4"
)

type (
	SampleEchoController struct {
		gpmu usecase.GetSampleUsecaseInterface
	}
	SampleEchoControllerInterface interface {
		ExecuteGetSample(ctx echo.Context) (int, any)
	}
)

func NewSampleEchoController(gpmu usecase.GetSampleUsecaseInterface) SampleEchoControllerInterface {
	return &SampleEchoController{gpmu: gpmu}
}

func (pc *SampleEchoController) ExecuteGetSample(ctx echo.Context) (int, any) {
	Sample, err := pc.gpmu.Execute()
	return response.GenerateEchoResponse(ctx, Sample, err)
}
