package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Задаем маршруты для обработки HTTP-запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to My Fitness App!")
	})

	// Запускаем HTTP-сервер на порту  8080
	port := ":8080"
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
