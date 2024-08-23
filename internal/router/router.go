package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/handler"
)

func SetupRoute(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.GET("/books", handler.ReadAllBooks)

		api.POST("/books", handler.CreateBook)

		api.GET("/books/:id", handler.ReadBook)

		api.PUT("/books/:id", handler.UpdateBook)

		api.DELETE("/books/:id", handler.DeleteBook)
	}
}
