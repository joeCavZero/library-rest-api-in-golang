package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/database"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/model"
)

func UpdateBook(ctx *gin.Context) {

	//---- getting the id from the url
	aux_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}
	id := uint32(aux_id)

	//---- creating a target book
	target_book := model.Book{}

	err = ctx.BindJSON(&target_book)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}

	err = target_book.Validate()
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}

	database.UpdateBookById(id, target_book)
}
