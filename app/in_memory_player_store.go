package main

// Initializes the store. For convenience
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := i.store[name]
	return score, ok
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
