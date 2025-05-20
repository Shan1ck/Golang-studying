package main

import (
	"asynchrony/library/MutexMap"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	m := MutexMap.NewStorage(map[string]float64{
		"Alex":  10.0,
		"Paul":  40.0,
		"Frank": 15.0,
	})
	m.Print()

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, key := range m.GetKeys() {
				time.Sleep(time.Millisecond * 10)
				m.IncreaseValue(key, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println()
	m.Print()
}
