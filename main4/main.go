package main

import (
	"fmt"
)

func main() {
	workDay := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	holiDays := workDay[5:7]
	workDay = workDay[:5]
	allDays := append(workDay, holiDays...)
	fmt.Println(holiDays)
	fmt.Println(workDay)
	fmt.Println(allDays)
}
