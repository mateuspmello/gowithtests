package contexto

import (
	"fmt"
	"net/http"
	"time"
)

//Store é uma loja
type Store interface {
	Fetch() string
	Cancel()
}

//SpyStore é a loja para espião
type SpyStore struct {
	response  string
	cancelled bool
}

//Fetch da Store
func (s *SpyStore) Fetch() string {
	time.Sleep(1000 * time.Millisecond)
	return s.response
}

//Cancel da Store
func (s *SpyStore) Cancel() {
	s.cancelled = true
}

//Server inicia o servidor
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}

	}
}
