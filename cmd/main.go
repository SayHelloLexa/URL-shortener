package main

import (
	"log"

	"github.com/sayhellolexa/url-short/internal/server"
)

func main() {
	s := server.NewServer()
	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
