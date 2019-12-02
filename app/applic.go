package app

import (
	"fmt"
	"net/http"
)

//PlayerStore é a interface que define os metodos
type PlayerStore interface {
	GetPlayerValue(player string) int
	RecordWin(player string)
}

//PlayerServer o servidor de player tem uma interface playerstore
type PlayerServer struct {
	store PlayerStore
}

//ServeHTTP serve para escrever no responsewriter
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {

	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerValue(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

//GetPlayerValue pega o score do player
func GetPlayerValue(player string) int {
	if player == "Pepper" {
		return 20
	}

	if player == "Floyd" {
		return 10
	}

	return 0
}

// func ListenAndServe(addr string, handler Handler) error
