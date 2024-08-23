package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/model"
)

func sendError(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, model.NewResponseError(err.Error()))
}
