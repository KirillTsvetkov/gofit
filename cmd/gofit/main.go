package main

import (
	"log"

	"github.com/KirillTsvetkov/gofit"
	"github.com/KirillTsvetkov/gofit/config"
	"github.com/KirillTsvetkov/gofit/handler"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	handeler := new(handler.Handler)
	srv := new(gofit.Server)
	if err := srv.Run(viper.GetString("port"), handeler.IniteRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
