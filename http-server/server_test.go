package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Matthew": 20,
			"Tom":     10,
		},
	}
	server := &PlayerServer{store: &store}

	t.Run("returns score of given username", func(t *testing.T) {
		username := "Matthew"

		req := newGetScoreRequest(username)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("returns score of another username", func(t *testing.T) {
		username := "Tom"

		req := newGetScoreRequest(username)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
	})
}

func newGetScoreRequest(playername string) *http.Request {
	url := fmt.Sprintf("/player/%s", playername)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
