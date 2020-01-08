package jfilho

import (
	"bufio"
	"io"
	"strings"
)

//CLI é a estrutura criada para o command line interface
type CLI struct {
	PlayerStore PlayerStore
	In          io.Reader
}

//PlayPoker inicia o jogo de Poker via CLI
func (c CLI) PlayPoker() {
	reader := bufio.NewScanner(c.In)
	reader.Scan()
	c.PlayerStore.RecordWin(extractWinner(reader.Text()))

}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
