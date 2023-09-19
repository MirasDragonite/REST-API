package main

import (
	"log"

	"rest/pkg/handlers"
	"rest/pkg/repository"
	"rest/pkg/service"

	structs "rest"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)
	srv := new(structs.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
