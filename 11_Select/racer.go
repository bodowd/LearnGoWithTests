package racer

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	// select lets you wait on multiple channels
	// the first one to send a value "wins" and the code underneath the case is executed
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}

}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	// sends a signal into the channel when http.Get(url) completes
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
