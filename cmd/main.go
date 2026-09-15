package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "server: ", log.LstdFlags)

	srv := server.New(logger)

	log.Println("Сервер запускается на порту :8080...")

	if err := srv.Start(); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
