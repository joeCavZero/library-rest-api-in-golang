package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joeCavZero/library-rest-api-in-golang/internal/model"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/util/logkit"
)

func CreateBook(ctx *gin.Context) {
	logger := logkit.New("CreateBook POST")
	var request model.BookResponse

	err := ctx.BindJSON(&request)
	if err != nil {
		responseError := model.ResponseError{
			Message: err.Error(),
		}

		ctx.JSON(http.StatusBadRequest, responseError)

		logger.Error(err)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		request,
	)

}
