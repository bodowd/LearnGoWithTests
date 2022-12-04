package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

// testing.TB is an interface so that you can call helper functions from a test
// or benchmark
func assertCorrectMessage(t testing.TB, got string, want string) {
	// this is needed to tell the test suite that this method is a helper
	// this will make the line number reported when a test fails be from our
	// function call rather than inside our test helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
