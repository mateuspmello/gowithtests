package balancear

import "testing"

func TestBalanco(t *testing.T) {

	mensagemDeErro := func(t *testing.T, want bool, got bool) {
		if want != got {
			t.Errorf("Esperava %v mas obteve %v", want, got)
		}
	}

	t.Run("Testando balanco correto", func(t *testing.T) {
		got := Balance("(a[0]+b[2c[6]]) {24+53}")
		want := true
		mensagemDeErro(t, want, got)
	})

	t.Run("Testando balanco correto", func(t *testing.T) {
		got := Balance("f(e(d))")
		want := true
		mensagemDeErro(t, want, got)
	})

	t.Run("Testando balanco correto", func(t *testing.T) {
		got := Balance("[()]{}([])")
		want := true
		mensagemDeErro(t, want, got)
	})

	t.Run("Testando balanco incorreto", func(t *testing.T) {
		got := Balance("([)]")
		want := false
		mensagemDeErro(t, want, got)
	})

	t.Run("Testando balanco incorreto", func(t *testing.T) {
		got := Balance("(c]")
		want := false
		mensagemDeErro(t, want, got)
	})

	t.Run("Testando balanco incorreto", func(t *testing.T) {
		got := Balance("((b)")
		want := false
		mensagemDeErro(t, want, got)
	})

	t.Run("Testando balanco incorreto", func(t *testing.T) {
		got := Balance("{(a[])")
		want := false
		mensagemDeErro(t, want, got)
	})

	t.Run("Testando balanco incorreto", func(t *testing.T) {
		got := Balance("([)]")
		want := false
		mensagemDeErro(t, want, got)
	})

	t.Run("Testando balanco incorreto", func(t *testing.T) {
		got := Balance(")(")
		want := false
		mensagemDeErro(t, want, got)
	})

	t.Run("Testando balanco em branco", func(t *testing.T) {
		got := Balance("")
		want := false
		mensagemDeErro(t, want, got)

	})
}
