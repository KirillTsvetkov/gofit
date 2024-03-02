package main

import (
	"log"

	"github.com/KirillTsvetkov/gofit/router"

	"github.com/KirillTsvetkov/gofit"
	"github.com/KirillTsvetkov/gofit/config"
	"github.com/KirillTsvetkov/gofit/repository"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	dbClient, _ := repository.NewMongoDBClient()

	rep := repository.NewRepository(dbClient)

	router := new(router.Router)
	srv := new(gofit.Server)
	if err := srv.Run(viper.GetString("port"), router.IniteRoutes(rep)); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
