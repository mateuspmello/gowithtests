package syncro

import "sync"

//Counter é um contador inteiro
type Counter struct {
	valor int
	mu    sync.Mutex
}

//Inc Add 1 em Counter
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

//Valor Add 1 em Counter
func (c *Counter) Valor() int {
	return c.valor
}
