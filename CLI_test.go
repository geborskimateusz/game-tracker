package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("Chris wins", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.playPoker()

		want := "Chris"

		assertPlayerWin(t, playerStore, want)
	})

}

func assertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("expected winner %q got %q", store.winCalls[0], winner)
	}
}
