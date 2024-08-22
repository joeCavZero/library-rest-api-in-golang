package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/database"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/model"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/util/logkit"
)

func ReadBook(ctx *gin.Context) {
	lk := logkit.New("ReadBook")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.NewResponseError("Invalid or missing id"),
		)
		return
	}

	new_book, err := database.GetBookById(uint32(id))
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			model.NewResponseError("Book not found"),
		)
		return
	}

	lk.Infof("Book found: %v", new_book)
	ctx.JSON(
		http.StatusOK,
		new_book,
	)

}
