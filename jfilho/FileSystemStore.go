package jfilho

import (
	"encoding/json"
	"io"
	"os"
	"sort"
)

//FileSystemPlayerStore sera a base de dados
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

//ReadWriteSeekTruncate é uma interface para ajudar a capturar requisitos
type ReadWriteSeekTruncate interface {
	io.ReadWriteSeeker
	Truncate(size int64) error
}

//NewFileSystemPlayerStore retorna a nova liga
func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}
}

//GetLeague obtem a liga
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})

	return f.league
}

//GetPlayerScore obtem o score de um jogador baseado no nome
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

//RecordWin adiciona um em Win do Player
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}

	f.database.Encode(f.league)
}
