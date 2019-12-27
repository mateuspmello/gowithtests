package jfilho

import (
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "123456")
	defer clean()

	tape := tape{file}
	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContens, _ := ioutil.ReadAll(file)

	got := string(newFileContens)
	want := "abc"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
