package mapas

import "testing"

var palavraNaoExiste = "Palavra não existe %s"
var esperavaRecebeu = "Esperava %s. Recebeu %s."

func TestAdiciona(t *testing.T) {

	t.Run("Teste de adicionar um elemento", func(t *testing.T) {
		dicionario := Dicionario{}
		palavra := "carro"
		definicao := "carro é um bem de consumo"
		dicionario.Adiciona(palavra, definicao)

		comparaDefinicao(t, dicionario, palavra, definicao)
	})
}

func comparaDefinicao(t *testing.T, d Dicionario, p string, de string) {
	t.Helper()
	got, err := d.Busca(p)
	want := de
	if err != nil {
		t.Fatalf(palavraNaoExiste, p)
	}

	if got != want {
		t.Errorf(esperavaRecebeu, want, got)
	}
}
