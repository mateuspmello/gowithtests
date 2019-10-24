package iteracao

import "testing"

func TestRepete(t *testing.T) {

	testeGeral := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Esperava %s e obteve %s", want, got)
		}
	}

	t.Run("Testando Repetição de 5 Caracteres", func(t *testing.T) {
		got := Repete("a", 3)
		want := "aaa"
		testeGeral(t, got, want)
	})
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repete("a", 3)
	}
}
