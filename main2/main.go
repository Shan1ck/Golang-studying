package main

import (
	"fmt"
)

var (
	A *int
	B = 32
)

func main() {
	A = &B
	*A = 52
	fmt.Println(*A, B)

}
