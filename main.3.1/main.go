package main

import (
	"fmt"

	"github.com/huandu/xstrings"
	"main3.1/task/city"
	"main3.1/task/digit"
)

func main() {
	fmt.Println(xstrings.Shuffle(city.City()))
	fmt.Println(digit.Digit())
}
