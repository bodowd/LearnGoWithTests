package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// Fprint is like fmt.Printf but instead takes a Writer type which it sends the string to, whereas Printf
	// defaults to stdout
	fmt.Fprintf(writer, "Hello, %s", name)
}

// http.ResponseWriter also implements io.Writer so this is why we could reuse the Greet function
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
