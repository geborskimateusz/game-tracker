package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: database,
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(playername string) {
	league := f.league
	player := f.league.Find(playername)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{Name: playername, Wins: 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
