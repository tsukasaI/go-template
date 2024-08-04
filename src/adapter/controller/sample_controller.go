package controller

import (
	"my-go-template/src/adapter/controller/response"
	"my-go-template/src/usecase"

	"github.com/gin-gonic/gin"
)

type (
	SampleController struct {
		gpmu usecase.GetSampleUsecaseInterface
	}
	SampleControllerInterface interface {
		ExecuteGetSample(ctx *gin.Context)
	}
)

func NewSampleController(gpmu usecase.GetSampleUsecaseInterface) SampleControllerInterface {
	return &SampleController{gpmu: gpmu}
}

func (pc *SampleController) ExecuteGetSample(ctx *gin.Context) {
	Sample, err := pc.gpmu.Execute("")
	response.GenerateResponse(ctx, Sample, err)
}
