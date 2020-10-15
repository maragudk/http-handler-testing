package main

import (
	"log"

	"http-handler-testing/server"
	"http-handler-testing/storage"
)

func main() {
	s := server.Server{Storer: &storage.Storer{}}
	if err := s.Start(); err != nil {
		log.Fatalln(err)
	}
}
