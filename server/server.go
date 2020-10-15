package server

import (
	"net/http"

	"github.com/go-chi/chi"

	"http-handler-testing/handlers"
	"http-handler-testing/storage"
)

type Server struct {
	Storer *storage.Storer
}

func (s *Server) Start() error {
	mux := chi.NewMux()

	handlers.PartyStarter(mux, s.Storer)

	return http.ListenAndServe(":8080", mux)
}
