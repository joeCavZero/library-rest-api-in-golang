package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/handler"
)

func SetupRoute(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.GET("/books", handler.ReadAllBooks)

		api.POST("/book", handler.CreateBook)

		api.GET("/book/:id", handler.ReadBook)

		api.PUT("/book/:id", handler.UpdateBook)

		api.DELETE("/book/:id", handler.DeleteBook)
	}
}
