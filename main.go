package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

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
		fmt.Println(line)
	}
	f.Close()
}
