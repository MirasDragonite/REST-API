package main

import (
	"log"

	"rest/pkg/handlers"

	structs "rest"
)

func main() {
	handlers := new(handlers.Handler)
	srv := new(structs.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
