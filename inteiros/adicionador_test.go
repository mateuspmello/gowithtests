package main

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {

	metodo := func(t *testing.T, got int, want int) {
		t.Helper()
		if got != want {
			t.Errorf("Esperava %d mas veio %d", want, got)
		}
	}

	t.Run("Testando uma soma simples", func(t *testing.T) {
		got := Adicionar(2, 2)
		want := 4
		metodo(t, got, want)
	})
}

func ExampleAdicionar() {
	soma := Adicionar(1, 5)
	fmt.Println(soma)
	// Output: 6
}
