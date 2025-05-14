package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)

	n := 44
	k := 45

	go fib(n)
	go fib(k)

	fmt.Printf("\rFibonachi(%d) = %d\n", n, fib(n))
	fmt.Printf("\rFibonachi(%d) = %d\n", k, fib(k))

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
