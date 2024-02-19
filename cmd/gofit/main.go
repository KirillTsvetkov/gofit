package main

import (
	"log"

	"github.com/KirillTsvetkov/gofit"
	"github.com/KirillTsvetkov/gofit/pkg/handler"
)

func main() {
	handeler := new(handler.Handler)
	srv := new(gofit.Server)
	if err := srv.Run("8080", handeler.IniteRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
