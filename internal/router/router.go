package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/handler"
)

func SetupRoute(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.GET("/books", func(ctx *gin.Context) {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"message": "HELOOOOOO",
				},
			)
		})

		api.POST("/books", handler.CreateBook)
	}
}
