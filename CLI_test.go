package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/geborskimateusz/game-tracker"
)

func TestCLI(t *testing.T) {
	var anySpyAlerter = &SpyBlindAlerter{}

	t.Run("Chris wins", func(t *testing.T) {
		stdin := strings.NewReader("Chris wins\n")
		stdout := &bytes.Buffer{}
		playerStore := &poker.StubPlayerStore{}
		game := &GameSpy{}

		cli := poker.NewCLI(stdin, stdout, game)

		cli.PlayPoker()

		want := "Chris"

		poker.AssertPlayerWin(t, playerStore, want)
	})

	t.Run("Cleo wins", func(t *testing.T) {
		stdin := strings.NewReader("Cleo wins\n")
		stdout := &bytes.Buffer{}
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, stdin, stdout, anySpyAlerter)

		cli.PlayPoker()

		want := "Cleo"

		poker.AssertPlayerWin(t, playerStore, want)
	})

}

type GameSpy struct {
	StartCalledWith  int
	FinishCalledWith string
}
