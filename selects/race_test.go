package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorrida(t *testing.T) {
	t.Run("teste de corrida", func(t *testing.T) {

		fastServer := makeDelayedServer(0 * time.Millisecond)

		slowServer := makeDelayedServer(20 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastURL
		got, _ := Corrida(slowURL, fastURL)

		if got != want {
			t.Fatalf("Esperava %v Recebeu %v", want, got)
		}
	})
	t.Run("teste de corrida - erro por tempo", func(t *testing.T) {

		fastServer := makeDelayedServer(12 * time.Second)

		slowServer := makeDelayedServer(20 * time.Second)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Corrida(fastURL, slowURL)
		if err != nil {
			t.Errorf("Timeout %v", err)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
