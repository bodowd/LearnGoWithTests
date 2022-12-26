package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()

	for i, player := range league {
		if player.Name == name {
			// needs to be this way instead of player.Wins++
			// when you range over a slice you are returned the current index
			// and a COPY of the element at that index
			// so player.Wins++ modifies a copy, not the underlying
			// slice that we are iterating on
			league[i].Wins++
		}
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
