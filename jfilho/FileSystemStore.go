package jfilho

import (
	"encoding/json"
	"io"
)

//FileSystemPlayerStore sera a base de dados
type FileSystemPlayerStore struct {
	db io.ReadWriteSeeker
}

//GetLeague obtem a liga (conjunto de jogadores)
func (f *FileSystemPlayerStore) GetLeague() League {
	f.db.Seek(0, 0)
	league, _ := NewLeague(f.db)
	return league
}

//GetPlayerScore obtem o score de um jogador baseado no nome
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var ret int
	liga := f.GetLeague()
	player := liga.Find(name)
	if player != nil {
		ret = player.Wins
	}
	return ret
}

//RecordWin adiciona um em Win do Player
func (f *FileSystemPlayerStore) RecordWin(name string) {
	liga := f.GetLeague()
	player := liga.Find(name)
	if player != nil {
		player.Wins++
	}

	f.db.Seek(0, 0)
	json.NewEncoder(f.db).Encode(liga)
}
