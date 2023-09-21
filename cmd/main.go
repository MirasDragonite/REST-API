package main

import (
	"log"

	"rest/pkg/handlers"
	"rest/pkg/repository"
	"rest/pkg/service"

	structs "rest"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewDB()
	if err  !=nil{
		log.Fatalf("failed to connect db %s",err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)
	srv := new(structs.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
