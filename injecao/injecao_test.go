package injecao

import (
	"bytes"
	"testing"
)

func TestSaudar(t *testing.T) {
	buffer := bytes.Buffer{}
	Saudar(&buffer, "Mateus")
	got := buffer.String()
	want := "Hello, Mateus"

	if got != want {
		t.Errorf("Esperava: %s. Recebeu: %s")
	}
}
