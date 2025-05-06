package main // Fixed package declaration

import (
	"fmt"
)

func main() {
	infBooks := make(map[string]map[string][]string)

	infBooks["Vasya"] = map[string][]string{
		"Books":     {"WAP", "Crime"},
		"magazines": {"Math everywhere"}}

	infBooks["Masha"] = map[string][]string{
		"Books":     {"WAP"},
		"magazines": {"Math everywhere", "New biology"}}

	infBooks["Den"] = map[string][]string{
		"Books":     {"WAP", "Crime", "1984"},
		"magazines": {}}

	// Print total readers
	fmt.Println("Total readers:", len(infBooks))

	// Calculate publications per reader
	for reader := range infBooks {
		total := 0
		// Iterate through publication categories
		total += len(infBooks[reader]["magazines"])
		fmt.Printf("%s has %d items\n", reader, total)
	}

}
