package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func (cli *CLI) PlayPoker() {
	// bufio.Scanner reads the input from the io.Reader
	reader := cli.in
	reader.Scan()
	cli.playerStore.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}
