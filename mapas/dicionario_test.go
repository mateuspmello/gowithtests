package mapas

import "testing"

const (
	palavraNaoExiste  = "Palavra não existe %s"
	palavraDuplicada  = "Tentou inserir palavra duplicada %s"
	esperavaRecebeu   = "Esperava %s. Recebeu %s."
	atualizaDefinicao = "Não foi atualizada a palavra %s"
	deletarDefinicao  = "Não foi deletada a palavra %s"
)

func TestAdiciona(t *testing.T) {

	t.Run("Teste de adicionar um elemento", func(t *testing.T) {
		dicionario := Dicionario{}
		// dicionario := Dicionario{"carro": "carro é um bem de consumo"}
		palavra := "carro"
		definicao := "carro é um bem de consumo"
		descricao, err := dicionario.Adiciona(palavra, definicao)
		if err != nil {
			t.Fatalf(palavraDuplicada, palavra)
		}

		comparaErro(t, descricao, definicao)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})

}

func TestAtualiza(t *testing.T) {
	t.Run("Atualiza definicao", func(t *testing.T) {
		dicionario := Dicionario{"carro": "carro é um bem de consumo"}
		palavra := "carro"
		novaDefinicao := "carro é um patrimônio"
		err := dicionario.Atualiza(palavra, novaDefinicao)
		if err != nil {
			t.Errorf(palavraNaoExiste, palavra)
		}

		if dicionario[palavra] != novaDefinicao {
			t.Fatalf(novaDefinicao, palavra)
		}
	})
}

func TestDeletar(t *testing.T) {
	t.Run("Deleta definicao", func(t *testing.T) {
		dicionario := Dicionario{"carro": "carro é um bem de consumo"}
		palavra := "carro"
		dicionario.Deletar(palavra)

		def, _ := dicionario.Busca(palavra)
		if def == "carro é um bem de consumo" {
			t.Fatalf("Não deletou %s", palavra)
		}

	})
}

func comparaErro(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf(esperavaRecebeu, want, got)
	}
}

func comparaDefinicao(t *testing.T, d Dicionario, p string, de string) {
	t.Helper()
	got, err := d.Busca(p)
	want := de
	if err != nil {
		t.Fatalf(palavraNaoExiste, p)
	}
	comparaErro(t, got, want)
}
