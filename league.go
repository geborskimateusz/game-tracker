package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(playername string) *Player {
	for i, p := range l {
		if p.Name == playername {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(rdr io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
