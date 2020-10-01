package main

import (
	"encoding/json"
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
	server := NewPlayerServer(&store)

	t.Run("returns score of given username", func(t *testing.T) {
		username := "Matthew"

		req := newGetScoreRequest(username)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseBody(t, res.Body.String(), "20")
		assertStatusCode(t, res.Code, http.StatusOK)
	})

	t.Run("returns score of another username", func(t *testing.T) {
		username := "Tom"

		req := newGetScoreRequest(username)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseBody(t, res.Body.String(), "10")
		assertStatusCode(t, res.Code, http.StatusOK)
	})

	t.Run("404 on missing player", func(t *testing.T) {
		username := "John"

		req := newGetScoreRequest(username)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Code
		want := http.StatusNotFound

		assertStatusCode(t, got, want)
	})
}

func TestStoreWin(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted POST", func(t *testing.T) {
		player := "Pepper"

		req := newPostWinRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("epeced player %q got %q", player, store.winCalls[0])
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		var got []Player

		err := json.NewDecoder(res.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", res.Body, err)
		}

		assertStatusCode(t, res.Code, http.StatusOK)
	})

}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected status %d, got %d", want, got)
	}
}

func newPostWinRequest(playername string) *http.Request {
	url := fmt.Sprintf("/player/%s", playername)
	res, _ := http.NewRequest(http.MethodPost, url, nil)
	return res
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
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
