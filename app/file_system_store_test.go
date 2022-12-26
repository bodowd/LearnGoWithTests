package main

import (
	"io"
	"os"
	"testing"
)

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database := `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`
		tmpfile, cleanDatabase := createTempFile(t, database)
		defer cleanDatabase()
		store := FileSystemPlayerStore{tmpfile}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`
		tmpfile, cleanDatabase := createTempFile(t, database)
		defer cleanDatabase()
		store := FileSystemPlayerStore{tmpfile}

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database := `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`
		tmpfile, cleanDatabase := createTempFile(t, database)
		defer cleanDatabase()
		store := FileSystemPlayerStore{tmpfile}

		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})

}
