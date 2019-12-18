package main

import (
	"log"
	"net/http"

	"github.com/mateuspmello/testes/jfilho"
)

func main() {
	server := &jfilho.PlayerServer{jfilho.NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
