package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestContagem(t *testing.T) {

	t.Run("Contagem para baixo", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpySleeper{}
		Contagem(buffer, spy)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("Esperava: %v Recebeu: %v", want, got)
		}

		if spy.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spy.Calls)
		}
	})
}
