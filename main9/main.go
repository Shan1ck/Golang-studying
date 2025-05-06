package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("data/in.txt")

	if err != nil {

		log.Fatalf("Couldn't open file. %v", err)

	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Couldn't close file.%v", err)
		}
	}()

	reader := bufio.NewScanner(file)
	countStr := 0

	for !reader.Scan() {
		if reader.Err() == nil {
			log.Fatal("File is empty", reader.Err())
		}
		log.Fatal("Error: couldn't read file")
	}

	for reader.Scan() {
		countStr++
	}

	fmt.Printf("Total strings: %d", countStr)
}
