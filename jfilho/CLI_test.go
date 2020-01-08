package jfilho

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("Testando para input Mateus", func(t *testing.T) {
		in := strings.NewReader("Mateus wins\n")
		playerStore := &StubPlayerStore{}
		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		if len(playerStore.winCalls) != 1 {
			t.Fatal("expected a win call but didn't get any")
		}

		assertPlayerWin(t, playerStore, "Mateus")
	})

	t.Run("Testando para input Joao", func(t *testing.T) {
		in := strings.NewReader("Joao wins\n")
		playerStore := &StubPlayerStore{}
		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		if len(playerStore.winCalls) != 1 {
			t.Fatal("expected a win call but didn't get any")
		}

		assertPlayerWin(t, playerStore, "Joao")
	})

}

func assertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], winner)
	}
}
