package handler

import (
	"github.com/gin-gonic/gin"
)

type responseError struct {
	Message string `json:"message"`
}

func sendError(ctx *gin.Context, code int, err error) {
	ctx.JSON(
		code,
		responseError{
			Message: err.Error(),
		},
	)
}
