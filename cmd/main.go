package main

import (
	"log"

	"github.com/sayhellolexa/url-short/internal/config"
	"github.com/sayhellolexa/url-short/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	s := server.NewServer(cfg)
	err = s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
