package server

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)

		go func() {
			// recall: send statement
			// store.Fetch() is sent to the `data` channel
			data <- store.Fetch()
		}()

		// recall: select lets you wait on multiple channels
		// we use select here to race the two asynchronous processes and we either
		// write a response or Cancel
		select {
		// recall: this is a receive expression
		// channel is on the right (`data` is the channel)
		// assigns a value received from a channel to a variable
		case d := <-data:
			fmt.Fprint(w, d)

			// context has a method Done() which returns a channel which gets sent
			// a signal when the context is "done" or "cancelled"
			// we want to listen to that signal and call store.Cancel if we get it
			// but ignore it if our Store manages to Fetch before it
			// recall: <- is a send statement
			// this says "if ctx.Done gets sent to the channel"
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
