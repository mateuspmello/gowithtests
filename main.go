package main

import (
	"log"
	"net/http"

	"github.com/mateuspmello/testes/app"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerValue(name string) int {
	return 123
}

func main() {
	server := &app.PlayerServer{&InMemoryPlayerStore{nil, nil}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
