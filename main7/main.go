package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file1, err := os.Open("06_task_in.txt")
	if err != nil {
		panic(err)
	}
	defer file1.Close()

	file2, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	scanner := bufio.NewScanner(file1)
	writer := bufio.NewWriter(file2)
	defer writer.Flush()
	i := 1
	for scanner.Scan() {
		line := fmt.Sprintf("%d. %s\n", i, scanner.Text())
		i++
		_, err := writer.WriteString(line)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
