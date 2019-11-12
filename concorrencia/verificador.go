package concorrencia

//VerificadorWebsite tipo de verificador de website
type VerificadorWebsite func(string) bool
type resultado struct {
	string
	bool
}

//VerificadorWebsites faz a verificacao de modo paralelo atraves de goroutines e channels
func VerificadorWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)
	canalResultado := make(chan resultado)

	for _, url := range urls {
		go func(u string) {
			canalResultado <- resultado{u, vw(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		resultado := <-canalResultado
		resultados[resultado.string] = resultado.bool
	}

	return resultados
}