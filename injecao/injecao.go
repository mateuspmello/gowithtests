package injecao

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

//Saudar printa o hello world
func Saudar(escritor *bytes.Buffer, nome string) {
	fmt.Fprintf(escritor, "Hello, %s", nome)
}

//SaudarW printa o hello world
func SaudarW(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Hello, %s", nome)
}

//MinhaSaudacao é o handler
func MinhaSaudacao(w http.ResponseWriter, r *http.Request) {
	SaudarW(w, "mundo")
}
