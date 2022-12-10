package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// select lets you wait on multiple channels
	// the first one to send a value "wins" and the code underneath the case is executed
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil

		// good to have a case where a signal is sent to select incase the code above
		// blocks forever because the channels you're listening on never return a value

		// time.After also returns a channel so we can use it here like so
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

}
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
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
