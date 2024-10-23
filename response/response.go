package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walnut/exception"
)

const (
	Ok                  = 1000
	InternalServerError = 1001
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(
		http.StatusOK,
		Response{
			Code:    200,
			Message: "success",
			Data:    data,
		},
	)
}

func Failed(ctx *gin.Context, err error) {
	httpCode := http.StatusInternalServerError
	val, ok := err.(*exception.ApiException)
	if ok {
		if val.HttpCode != 0 {
			httpCode = val.HttpCode
		}
	} else {
		val = exception.ErrServerInternal(err.Error())
	}

	ctx.JSON(
		httpCode,
		Response{
			Code:    val.Code,
			Message: val.Error(),
			Data:    nil,
		},
	)
	ctx.Abort()
}
