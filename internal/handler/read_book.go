package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/database"
)

func ReadBook(ctx *gin.Context) {
	//---- getting the id from the url
	aux_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}
	id := uint32(aux_id)

	//---- getting the book from the database
	new_book, err := database.GetBookById(id)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		new_book,
	)

}
