package mapas

//Dicionario de palavras
type Dicionario map[string]string

//ErrDicionario erros do dicionario
type ErrDicionario string

var (
	errNaoEncontrado    = ErrDicionario("não foi possível encontrar a palavra que você procura")
	errDuplicado        = ErrDicionario("palavra %s já existe. Descrição: %s")
	errPalavraNaoExiste = ErrDicionario("palavra não existe")
)

//Busca uma palavra
func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return palavra, errNaoEncontrado
	}
	return definicao, nil
}

//Adiciona uma palavra
func (d Dicionario) Adiciona(palavra string, descricao string) (string, error) {
	_, existe := d[palavra]
	if existe {
		return d[palavra], errDuplicado
	}
	d[palavra] = descricao
	return descricao, nil
}

func (err ErrDicionario) Error() string {
	return string(err)
}

//Atualiza faz a atualização da definição
func (d Dicionario) Atualiza(palavra string, novaDefinicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case errNaoEncontrado:
		return errPalavraNaoExiste
	case nil:
		d[palavra] = novaDefinicao
	}
	return nil
}

//Deletar faz a remoção do elemento do map
func (d Dicionario) Deletar(palavra string) {
	delete(d, palavra)
}
