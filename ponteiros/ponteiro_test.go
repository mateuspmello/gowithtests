package ponteiros

import "testing"

func TestPonteiro(t *testing.T) {

	mensagem := func(t *testing.T, c Carteira, want Bitcoin) {
		t.Helper()
		got := c.Saldo()
		if want != got {
			t.Errorf("Esperava %v recebeu %v", want, got)
		}
	}

	t.Run("Testando um deposito na conta", func(t *testing.T) {

		carteira := Carteira{saldo: 5}
		carteira.Depositar(Bitcoin(5))

		want := Bitcoin(10)

		mensagem(t, carteira, want)
	})

	t.Run("Testando uma retirada positiva", func(t *testing.T) {

		carteira := Carteira{saldo: 5}
		carteira.Retirar(Bitcoin(4))

		want := Bitcoin(1)
		mensagem(t, carteira, want)
	})

	t.Run("Testando uma retirada negativa", func(t *testing.T) {

		carteira := Carteira{saldo: 5, chequeEspecial: true}
		carteira.Retirar(Bitcoin(6))

		want := Bitcoin(-1)
		mensagem(t, carteira, want)
	})

	t.Run("Testando uma retirada negativa sem cheque", func(t *testing.T) {

		carteira := Carteira{saldo: 5, chequeEspecial: false}
		err := carteira.Retirar(Bitcoin(6))

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Testando uma retirada negativa sem setar chequeEspecial", func(t *testing.T) {

		carteira := Carteira{saldo: 5}
		carteira.Retirar(Bitcoin(6))

		want := Bitcoin(5)

		mensagem(t, carteira, want)
	})
}
