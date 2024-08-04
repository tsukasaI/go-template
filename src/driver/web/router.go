package web

import (
	"my-go-template/src/di"

	"github.com/gin-gonic/gin"
)

type Engine struct{ Engine *gin.Engine }

func NewEngine(e *gin.Engine) Engine {
	return Engine{e}
}

func (e *Engine) SetupRouter() {
	v1Router := e.Engine.Group("/v1")

	sampleController := di.GetSampleController()
	{
		v1Router.GET("/samples", sampleController.ExecuteGetSample)
	}

}
