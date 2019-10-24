package array

//Somar soma um array de int
func Somar(a []int) int {
	resultado := 0
	for _, i := range a {
		resultado += i
	}

	return resultado
}

//SomaTudo soma uma quantidade indefinida de slices e salva o resultado em outro slice
func SomaTudo(slices ...[]int) []int {
	var retorno []int
	for _, s := range slices {
		retorno = append(retorno, Somar(s))
	}
	return retorno
}
