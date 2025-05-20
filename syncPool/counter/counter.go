package counter

import (
	"fmt"
	"sync"
)

type Counter struct {
	Data []int
	A  int
	B  int
	mu sync.RWMutex
}

func NewCounter(initA, initB int) *Counter {
	return &Counter{
		Data: make([]int, 100),
		A: initA,
		B: initB,
	}
}

func (c *Counter) IncreaseValue() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.A += 1
	c.B += 1
}

func (c *Counter) Print() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	fmt.Println(c.A)
	fmt.Println(c.B)

}
