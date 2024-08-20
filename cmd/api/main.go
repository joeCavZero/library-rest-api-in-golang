package main

import (
	"github.com/joeCavZero/library-rest-api-in-golang/config"
	"github.com/joeCavZero/library-rest-api-in-golang/internal/server"
)

func main() {
	err := config.Initialize()
	if err != nil {
		panic(err)
	}

	server.StartServer()
}
