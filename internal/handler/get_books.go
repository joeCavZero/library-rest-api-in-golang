package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/database"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/util/logkit"
)

func ReadBooks(ctx *gin.Context) {
	lk := logkit.New("ReadBooks GET")
	books, err := database.GetBooks()
	if err != nil {
		lk.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, books)

}
