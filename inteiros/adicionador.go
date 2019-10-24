package main

import "fmt"

//Adicionar soma uma quantidade x de variaveis inteiras
func Adicionar(nuns ...int) int {

	retorno := 0
	for _, n := range nuns {
		retorno += n
	}

	return retorno
}

func main() {

	fmt.Printf("Resultado da soma: %d \n\r", Adicionar(1, 2, 3))
}
