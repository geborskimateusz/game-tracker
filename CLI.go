package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.store.RecordWin(extractWinner(userInput))
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		store: store,
		in:    bufio.NewScanner(in),
	}
}

func extractWinner(winnerCommand string) string {
	return strings.Replace(winnerCommand, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
