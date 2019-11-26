package syncro

import (
	"sync"
	"testing"
)

func NewCounter() *Counter {
	return &Counter{}
}

func TestCount(t *testing.T) {

	t.Run("Testando o contador normal", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 5)

	})

	t.Run("Testando o contador com concorrencia", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)

		}
		wg.Wait()

		assertCounter(t, &counter, 1000)

	})

}

func assertCounter(t *testing.T, c *Counter, want int) {
	t.Helper()
	if want != c.Valor() {
		t.Errorf("Esperava %v  Obteve %v", want, c.Valor())
	}
}
