package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/joeCavZero/library-rest-api-in-golang/internal/controller/database"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/model"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/util/logkit"
)

var databaseController = database.NewDatabaseController()

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

	book_response, err := databaseController.CreateBook(new_book)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err)
		lk.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, book_response)

}

func DeleteBook(ctx *gin.Context) {
	aux_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}
	id := uint32(aux_id)

	err = databaseController.DeleteBookById(id)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func ReadAllBooks(ctx *gin.Context) {
	lk := logkit.New("ReadBooks GET")
	books, err := databaseController.GetBooks()
	if err != nil {
		lk.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, books)

}

func ReadBook(ctx *gin.Context) {
	//---- getting the id from the url
	aux_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err)
		return
	}
	id := uint32(aux_id)

	//---- getting the book from the database
	new_book, err := databaseController.GetBookById(id)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		new_book,
	)

}

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

	databaseController.UpdateBookById(id, target_book)
}
