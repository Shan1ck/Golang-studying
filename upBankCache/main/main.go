package main

import (
	"BankCache/cache"
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	k1   = "key2"
	step = 7
	rt   = 4
	ts   = 10
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	c := cache.New()
	taskPool := make(chan func(), ts*2)

	for i := 0; i < rt; i++ {
		go func() {
			for task := range taskPool {
				task()
				wg.Done()
			}
		}()
	}

	wg.Add(ts * 2)

	for i := 0; i < ts; i++ {
		taskPool <- func() {
			c.Increase(k1, step)
		}
	}

	for i := 0; i < ts; i++ {
		i := i
		taskPool <- func() {
			c.Set(k1, step*i)
		}
	}

	go func() {
		wg.Wait()
		close(taskPool)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Time over")
		return
	case <-taskPool:
		fmt.Println("Final value:", c.Get(k1))
	}
}
