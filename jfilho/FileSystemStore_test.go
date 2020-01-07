package jfilho

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/ league form reader", func(t *testing.T) {
		database, cleanDatabase := createDatabase(t)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()
		/*sem problemas para executar duas vezes pq esta usando o ReadSeek que retorna
		* o reader para o começo novamente (f.db.Seek(0, 0))
		 */
		// got = store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)
	})

	t.Run("/ get Wins from Players", func(t *testing.T) {
		database, cleanDatabase := createDatabase(t)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("/ adicionar registro", func(t *testing.T) {
		database, cleanDatabase := createDatabase(t)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")
		want := 34
		got := store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, want)
	})

	t.Run("/store win for new players", func(t *testing.T) {
		database, cleanDatabase := createDatabase(t)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")
		want := 1
		got := store.GetPlayerScore("Pepper")

		assertScoreEquals(t, got, want)
	})

	t.Run("sorted league", func(t *testing.T) {
		database, cleanDatabase := createDatabase(t)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		//executando de novo
		got = store.GetLeague()
		assertLeague(t, got, want)

	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removefile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removefile
}

func createDatabase(t *testing.T) (*os.File, func()) {
	return createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
