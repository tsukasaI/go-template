package response

import (
	customErrors "my-go-template/src/core/errors"

	"github.com/gin-gonic/gin"
)

func GenerateResponse(ctx *gin.Context, result any, err customErrors.ErrInterface) {
	if err != nil {
		ctx.JSON(err.GetStatusCode(), NewErrorResponse(ctx.Request.Header.Get("X-Bond-Request-Id"), err))
		return
	}
	ctx.JSON(200, NewSuccessReponse(ctx.Request.Header.Get("X-Bond-Request-Id"), result))
}
