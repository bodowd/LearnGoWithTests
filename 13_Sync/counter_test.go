package counter

import (
	"sync"
	"testing"
)

/*
When to use locks over channels and goroutines?
- Use channels when passing ownership of data
- Use mutexes for managing state
*/

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		// Need this test to check if multiple goroutines are trying to mutate the value
		// of the counter at the same time
		wantedCount := 1000
		counter := NewCounter()

		// sync.WaitGroup is a way of synchronising concurrent processes
		// waits for a collection of goroutines to finish
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		// by waiting for wg.Wait() to finish before making our assertions, we
		// be sure all of our goroutines have attempted to Inc the Counter
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

// use this constructer which shows readers of your API that it would be better
// not to initialize the type yourself
func NewCounter() *Counter {
	// recall: `&` gives you the address of that bit of memory - i.e. 0xc42...
	return &Counter{}
}

// want to pass in a pointer to our Counter because otherwise, it will try to
// create a copy of the mutex
func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
