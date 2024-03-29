package poker

import (
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("get player score", func(t *testing.T) {
		database := `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`
		tmpfile, cleanDatabase := createTempFile(t, database)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(tmpfile)
		assertNoError(t, err)

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
		store, err := NewFileSystemPlayerStore(tmpfile)
		assertNoError(t, err)

		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database := `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`
		tmpfile, cleanDatabase := createTempFile(t, database)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(tmpfile)
		assertNoError(t, err)

		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database := `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`
		tmpfile, cleanDatabase := createTempFile(t, database)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(tmpfile)
		assertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

}
