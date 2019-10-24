package ponteiros

import "errors"

//Bitcoin é a moeda
type Bitcoin float64

//Carteira do bitcoin
type Carteira struct {
	saldo          Bitcoin
	chequeEspecial bool
}

//Depositar deposita um valor
func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

//Saldo infoma o saldo da carteira
func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

//Retirar deposita um valor
func (c *Carteira) Retirar(valor Bitcoin) error {

	if (c.saldo < valor && c.chequeEspecial) || c.saldo >= valor {
		c.saldo -= valor
	} else {
		return errors.New("Sem saldo")
	}
	return nil
}
