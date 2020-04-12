package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const defaultFile = "people.csv"

type Person struct {
	name string
	age  int
	city string
}

func main() {
	args := os.Args
	// fmt.Println(args)

	inputFile := defaultFile
	if len(args) > 1 {
		inputFile = args[1]
	}

	f, err := os.Open(inputFile)
	if err != nil {
		handleError(err)
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		handleError(err)
	}
	parseInput(lines)
	f.Close()
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func parseInput(lines [][]string) {
	for _, line := range lines {
		// fmt.Println(line)
		a, _ := strconv.Atoi(line[1])
		p := Person{
			name: line[0],
			age:  a,
			city: line[2],
		}
		fmt.Println(p)
	}
}
