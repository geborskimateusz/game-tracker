package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removefile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removefile
}

func TestFileSystemStore(t *testing.T) {

	t.Run("league from reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},	
			{"Name": "Tom", "Wins": 1}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Tom", 1},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},	
			{"Name": "Tom", "Wins": 1}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Tom")

		want := 1

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},	
			{"Name": "Tom", "Wins": 100}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Tom")

		got := store.GetPlayerScore("Tom")

		want := 101

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},	
			{"Name": "Tom", "Wins": 100}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Jenny")

		got := store.GetPlayerScore("Jenny")

		want := 1

		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})

	t.Run("retuns sorted league", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Randy", "Wins": 10},	
			{"Name": "Alice", "Wins": 111}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Alice", 111},
			{"Randy", 10},
		}

		assertLeague(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect error gt one, %v", err)
	}
}
