package main

import (
	"fmt"
)

func main() {
	b := []string{"1", "2", "3", "4"}
	go contains(b, "5")
	go getMax(1234, 2, 5, 7, 0, -1, 10, -100)
	fmt.Println(contains(b, "5"))
	fmt.Println(getMax(1234, 2, 5, 7, 0, -1, 10, -100))
}

func contains(slc []string, x string) bool {
	for _, i := range slc {
		if i == x {
			return true
		}
	}
	return false
}

func getMax(a ...int) int {
	aMax := 0
	for _, j := range a {
		if j > aMax {
			aMax = j
		}
	}
	return aMax
}
