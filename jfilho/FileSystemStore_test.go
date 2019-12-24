package jfilho

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/ league form reader", func(t *testing.T) {
		database, cleanDatabase := createDatabase(t)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()
		/*sem problemas para executar duas vezes pq esta usando o ReadSeek que retorna
		* o reader para o começo novamente (f.db.Seek(0, 0))
		 */
		// got = store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	})

	t.Run("/ get Wins from Players", func(t *testing.T) {
		database, cleanDatabase := createDatabase(t)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("/ adicionar registro", func(t *testing.T) {
		database, cleanDatabase := createDatabase(t)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		store.RecordWin("Chris")
		want := 34
		got := store.GetPlayerScore("Chris")

		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
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

func createDatabase(t *testing.T) (io.ReadWriteSeeker, func()) {
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
