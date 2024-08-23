package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/database"
)

func DeleteBook(ctx *gin.Context) {
	aux_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}
	id := uint32(aux_id)

	err = database.DeleteBookById(id)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
