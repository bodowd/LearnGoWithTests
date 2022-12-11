package counter

import "sync"

type Counter struct {
	// you  may see examples where they just put sync.Mutex straight instead
	// of mu sync.Mutex
	// this is bad because embedding types means the methods of that type
	// becomes part of the public interface and you often will not want that
	// this can cause unnecessary coupling
	// could allow counter.Lock() which should not be available to outsiders
	// calling Counter
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {

	// any goroutine calling Inc will acquire the lock on Counter if they are first
	// all other goroutines will have to wait for it to be Unlocked before getting
	// access
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++

}

func (c *Counter) Value() int {
	return c.value
}
