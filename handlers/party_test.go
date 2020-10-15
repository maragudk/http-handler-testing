package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

type partyStarterRepoMock struct {
	err error
}

func (m *partyStarterRepoMock) StartParty(id string) error {
	return m.err
}

func TestPartyStarter(t *testing.T) {
	t.Run("sends bad gateway on start party error", func(t *testing.T) {
		mux := chi.NewMux()
		PartyStarter(mux, &partyStarterRepoMock{err: errors.New("no snacks")})

		req := httptest.NewRequest(http.MethodPost, "/start/123", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		res := rec.Result()
		if res.StatusCode != http.StatusBadGateway {
			t.FailNow()
		}
	})

	t.Run("sends accepted on start party success", func(t *testing.T) {
		mux := chi.NewMux()
		PartyStarter(mux, &partyStarterRepoMock{})

		req := httptest.NewRequest(http.MethodPost, "/start/123", nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		res := rec.Result()
		if res.StatusCode != http.StatusAccepted {
			t.FailNow()
		}
	})
}
