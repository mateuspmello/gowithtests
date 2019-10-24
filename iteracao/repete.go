package iteracao

//Repete repete o caracter 5 vezes
func Repete(a string, qtd int) string {
	for i := 1; i < qtd; i++ {
		a += "a"
	}
	return a
}
