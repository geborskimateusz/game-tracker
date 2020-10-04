package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var score int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			score = player.Wins
			break
		}
	}

	return score
}

func (f *FileSystemPlayerStore) RecordWin(playername string) {
	league := f.GetLeague()

	for i, player := range league {
		if player.Name == playername {
			league[i].Wins++
		}
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
