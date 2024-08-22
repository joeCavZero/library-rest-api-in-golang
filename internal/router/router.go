package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/handler"
)

func SetupRoute(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.GET("/books", handler.ReadBooks)

		api.POST("/books", handler.CreateBook)
	}
}
