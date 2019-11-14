package selects

import (
	"fmt"
	"net/http"
	"time"
)

//Corrida faz a corrida
func Corrida(urlUm string, urlDois string) (winner string, error error) {
	select {
	case <-ping(urlUm):
		return urlUm, nil
	case <-ping(urlDois):
		return urlDois, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlUm, urlDois)
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()
	return ch
}
