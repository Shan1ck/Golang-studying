package main_test

import (
	"sync"
	"syncPool/counter"
	"testing"
)

func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			b.StopTimer()
			_ = counter.NewCounter(0, 2)
			b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return counter.NewCounter(0, 2)
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			c := pool.Get().(*counter.Counter)
			_ = c
			pool.Put(c)
		}
	}
}
