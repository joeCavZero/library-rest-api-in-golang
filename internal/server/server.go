package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joeCavZero/library-rest-api-in-golang/config"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/router"
)

func StartServer() {

	engine := gin.Default()

	router.SetupRoute(engine)

	engine.Run(
		fmt.Sprintf(":%d", config.PORT),
	)

}
