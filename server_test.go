package poker 

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
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

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		req := newLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := getLeagueFromResponse(t, res.Body)
		assertStatusCode(t, res.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentTypeAppJSON(t, res)
	})

}

func assertLeague(t *testing.T, got, want []Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func assertContentTypeAppJSON(t *testing.T, got *httptest.ResponseRecorder) {
	t.Helper()
	if got.Result().Header.Get("content-type") != "application/json" {
		t.Errorf("response did not hae content-type of appllication/json, got %v", got.Result().Header)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected status %d, got %d", want, got)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server")
	}
	return
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

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}
