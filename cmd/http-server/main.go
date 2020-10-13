package main

import (
	"log"
	"net/http"
	"os"
)

const dbFilename = "game.db.json"

func main() {

	// The 2nd argument to os.OpenFile lets you define the permissions
	// for opening the file, in our case O_RDWR means we want to read
	// and write and os.O_CREATE means create the file if it doesn't exist.
	// The 3rd argument means sets permissions for the file, in our case,
	// all users can read and write the file.
	db, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 6666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFilename, err)
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
