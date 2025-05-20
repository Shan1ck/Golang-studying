package main

import (
	"BankCache/cache"
	"fmt"
	"time"
)

const (
	k1   = "key1"
	step = 7
	rt   = 4
	ts   = 10
)

func main() {
	c := cache.New()
	taskPool := make(chan func(), ts*2)
	done := make(chan struct{}, ts*2)

	for i := 0; i < rt; i++ {
		go func() {
			for task := range taskPool {
				task()
				done <- struct{}{}
			}
		}()
	}

	for i := 0; i < ts; i++ {
		taskPool <- func() {
			c.Increase(k1, step)
			time.Sleep(time.Millisecond * 100)
		}
	}

	for i := 0; i < ts; i++ {
		i := i
		taskPool <- func() {
			c.Set(k1, step*i)
			time.Sleep(time.Millisecond * 100)
		}
	}

	for i := 0; i < ts*2; i++ {
		<-done
	}

	close(taskPool)
	fmt.Println("Final value:", c.Get(k1))
}
