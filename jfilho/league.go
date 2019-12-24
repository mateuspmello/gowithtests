package jfilho

import (
	"encoding/json"
	"fmt"
	"io"
)

//League é um array de Player
type League []Player

//NewLeague cria uma nova liga
func NewLeague(db io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(db).Decode(&league)
	if err != nil {
		err = fmt.Errorf("Error ao fazer o parsing: %v", err)
	}
	return league, err
}

//Find encontra o Player com base no nome
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
