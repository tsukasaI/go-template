package response

import (
	customErrors "my-go-template/src/core/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func GenerateResponse(ctx *gin.Context, result any, err customErrors.ErrInterface) {
	if err != nil {
		ctx.JSON(err.GetStatusCode(), NewErrorResponse(ctx.Request.Header.Get("X-Request-Id"), err))
		return
	}
	ctx.JSON(http.StatusOK, NewSuccessReponse(ctx.Request.Header.Get("X-Request-Id"), result))
}

func GenerateEchoResponse(ctx echo.Context, result any, err customErrors.ErrInterface) (int, any) {
	if err != nil {
		return NewEchoErrorResponse(ctx.Request().Header.Get("X-Request-Id"), err)
	}
	return http.StatusOK, NewSuccessReponse(ctx.Request().Header.Get("X-Request-Id"), result)
}
