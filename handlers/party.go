package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

type partyStarterRepo interface {
	StartParty(id string) error
}

func PartyStarter(mux chi.Router, s partyStarterRepo) {
	mux.Post("/start/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := s.StartParty(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})
}
