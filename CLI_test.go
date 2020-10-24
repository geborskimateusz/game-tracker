package poker_test

import (
	"strings"
	"testing"

	poker "github.com/geborskimateusz/game-tracker"
)

func TestCLI(t *testing.T) {

	t.Run("Chris wins", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Chris"

		poker.AssertPlayerWin(t, playerStore, want)
	})

	t.Run("Cleo wins", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Cleo"

		poker.AssertPlayerWin(t, playerStore, want)
	})

}
