package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("parse error:", r)
		}
	}()

	input, err := os.Open("data/input.txt")
	if err != nil {
		defer panic("file is not exist")
	}
	defer input.Close()

	output, err := os.Create("data/output.txt")
	if err != nil {
		defer panic("file is not exist")
	}
	defer output.Close()

	reader := bufio.NewScanner(input)

	writer := bufio.NewWriter(output)
	defer writer.Flush()

	rowsNum := 1

	for lineNum := 1; reader.Scan(); lineNum++ {
		line := reader.Text()
		data := strings.Split(line, "|")
		if len(data) != 3 {
			defer panic((fmt.Sprintf("empty field on string %d", rowsNum)))
		}
		name := strings.TrimSpace(data[0])
		address := strings.TrimSpace(data[1])
		city := strings.TrimSpace(data[2])

		if name == "" || address == "" || city == "" {
			defer panic(fmt.Sprintf("empty field on string %d", rowsNum))
		}
		outputInf := fmt.Sprintf(
			"Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n",
			rowsNum,
			name,
			address,
			city,
		)

		_, err = writer.WriteString(outputInf)
		if err != nil {
			panic("Writing of file is broken")
		}
		rowsNum++
	}

}
