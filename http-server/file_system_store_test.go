package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league from reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},	
			{"Name": "Tom", "Wins": 1}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Tom", 1},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},	
			{"Name": "Tom", "Wins": 1}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Tom")

		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
