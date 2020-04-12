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
	if len(args) >= 2 {
		inputFile = args[1]
	}

	cmds := []string{}
	if len(args) >= 3 {
		cmds = args[2:]
	}

	f, err := os.Open(inputFile)
	if err != nil {
		handleError(err)
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		handleError(err)
	}
	people := parseInput(lines)
	defer f.Close()

	// fmt.Println(people)
	find(people, cmds)
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func parseInput(lines [][]string) []Person {
	people := []Person{}
	for _, line := range lines {
		fmt.Println(line)
		a, _ := strconv.Atoi(line[1])
		p := Person{
			name: line[0],
			age:  a,
			city: line[2],
		}
		people = append(people, p)
	}
	return people
}

func find(people []Person, cmds []string) {
	cmd := cmds[0]
	switch cmd {
	case "oldest":
		fmt.Println("Looking for the oldest")
		person := oldest(people)
		fmt.Println("The oldest person is ", person)
	case "youngest":
		fmt.Println("Looking for the youngest")
	case "total":
		fmt.Println("Looking for the total number of people")
	case "city":
		fmt.Println(fmt.Sprintf("Looking for city %s", cmds[1]))
	case "name":
		fmt.Println(fmt.Sprintf("Looking for name %s", cmds[1]))
	default:
		fmt.Println("Unable to find", cmd)
	}
}

func oldest(people []Person) Person {
	oldest := 0
	oldestPerson := Person{}
	for _, person := range people {
		if person.age > oldest {
			oldest = person.age
			oldestPerson = person
		}
	}
	return oldestPerson
}
