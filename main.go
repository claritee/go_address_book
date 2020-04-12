package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	name string
	age  int
	city string
}

func main() {
	f, err := os.Open("people.csv")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

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
	f.Close()
}
