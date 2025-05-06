package main

import (
	"fmt"
	"math"
)

var (
	R *float64
	C = 32.00
)

func main() {
	Radius := C / (2 * math.Pi)
	R = &Radius
	S := math.Pi * (*R) * (*R)
	*R = math.Round(*R*100) / 100
	S = math.Round(S*100) / 100
	fmt.Println(*R, S)
}
