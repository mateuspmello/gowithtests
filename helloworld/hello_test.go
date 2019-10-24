package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper() // serve para pegar a linha de onde o metodo assertCorrectMessage foi chamado
		if got != want {
			t.Errorf("got '%s' but want '%s'", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Mateus")
		want := "Hello, Mateus"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello World if a empty string supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people in Portuguese", func(t *testing.T) {
		got := Hello("Mateus", "Portuguese")
		want := "Olá, Mateus"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Mateus", "Spanish")
		want := "Hola, Mateus"

		assertCorrectMessage(t, got, want)
	})
}
