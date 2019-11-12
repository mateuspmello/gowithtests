package mocking

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

const palavraFinal = "Go!"
const contagemRegressiva = 3

//Contagem Faz a contagem regressiva
func Contagem(saida io.Writer, sleeper Sleeper) {

	for i := contagemRegressiva; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := contagemRegressiva; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}
	sleeper.Sleep()
	fmt.Fprint(saida, palavraFinal)
}
