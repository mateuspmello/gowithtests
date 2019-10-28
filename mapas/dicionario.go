package mapas

import "errors"

//Dicionario de palavras
type Dicionario map[string]string

//ErrNaoEncontrado é a mensagem de erro não encontrado
var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra que você procura")

//Busca uma palavra
func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return palavra, ErrNaoEncontrado
	}
	return definicao, nil
}

//Adiciona uma palavra
func (d Dicionario) Adiciona(palavra string, descricao string) (string, error) {
	d[palavra] = descricao
	return descricao, nil
}
