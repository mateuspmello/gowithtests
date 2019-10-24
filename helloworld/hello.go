package main

import "fmt"

const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const spanishHelloPrefix = "Hola, "
const spanishLang = "Spanish"
const portugueseLang = "Portuguese"

//Hello retorna um hello world
func Hello(nome string, lingua ...string) string {
	ling := ""
	for _, v := range lingua {
		ling = v
	}
	if nome == "" {
		return englishHelloPrefix + "World"
	}

	if ling == portugueseLang {
		return portugueseHelloPrefix + nome
	}
	if ling == spanishLang {
		return spanishHelloPrefix + nome
	}

	return englishHelloPrefix + nome
}

func main() {
	fmt.Println(Hello("World"))
}
