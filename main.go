package main

import (
	"os"
	"time"

	"github.com/mateuspmello/testes/mocking"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// func Greet(writer io.Writer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "world")
// }

func main() {
	// http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	sleeper := &DefaultSleeper{}
	mocking.Contagem(os.Stdout, sleeper)
}
