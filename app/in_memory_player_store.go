package poker

import "sync"

// Initializes the store. For convenience
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{},
		sync.RWMutex{},
	}

}

type InMemoryPlayerStore struct {
	store map[string]int
	// a mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	score, ok := i.store[name]
	return score, ok
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
