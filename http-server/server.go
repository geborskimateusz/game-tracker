package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/player/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// func getUserScore(username string) string {
// 	switch username {
// 	case "Matthew":
// 		return "20"
// 	case "Tom":
// 		return "10"
// 	}

// 	return ""
// }
