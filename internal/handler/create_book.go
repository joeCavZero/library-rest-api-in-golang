package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/joeCavZero/library-rest-api-in-golang/internal/database"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/model"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/util/logkit"
)

func CreateBook(ctx *gin.Context) {
	lk := logkit.New("CreateBook POST")
	var new_book model.Book

	err := ctx.BindJSON(&new_book)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		lk.Error(err)
		return
	}
	lk.Infof("new_book: %v", new_book)

	err = new_book.Validate()
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		lk.Error(err)
		return
	}

	book_response, err := database.CreateBook(new_book)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err)
		lk.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, book_response)

}
