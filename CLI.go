package poker

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter number of players: "

type CLI struct {
	store   PlayerStore
	in      *bufio.Scanner
	out     io.Writer
	alerter BlindAlerter
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	cli.scheduleBlindAlerts()
	userInput := cli.readLine()
	cli.store.RecordWin(extractWinner(userInput))
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		store:   store,
		in:      bufio.NewScanner(in),
		out:     out,
		alerter: alerter,
	}
}

func extractWinner(winnerCommand string) string {
	return strings.Replace(winnerCommand, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
