package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"github.com/claritee/go_address_book/person"
)

const defaultFile = "people.csv"

// type Person struct {
// 	name string
// 	age  int
// 	city string
// }

type PersonNotFoundError struct{}

func (e *PersonNotFoundError) Error() string {
	return "Person not found"
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
		person := youngest(people)
		fmt.Println("The youngest person is ", person)
	case "total":
		fmt.Println("Looking for the total number of people")
		fmt.Println("Total number of people ", total(people))
	case "city":
		fmt.Println(fmt.Sprintf("Looking for city %s", cmds[1]))
		person, err := city(people, cmds[1])
		if err != nil {
			handleError(err)
		}
		fmt.Println("Looking for the person with city ", person)
	case "name":
		fmt.Println(fmt.Sprintf("Looking for name %s", cmds[1]))
		person, err := name(people, cmds[1])
		if err != nil {
			handleError(err)
		}
		fmt.Println("Looking for someone called ", person)
		// fallthrough - this is used to pass control flow to th next case
	default:
		fmt.Println("Unable to find", cmd)
	}
}

func name(people []Person, name string) (Person, error) {
	for _, person := range people {
		if person.name == name {
			return person, nil
		}
	}
	return Person{}, &PersonNotFoundError{}
}

func city(people []Person, city string) (Person, error) {
	for _, person := range people {
		if person.city == city {
			return person, nil
		}
	}
	return Person{}, &PersonNotFoundError{}
}

func total(people []Person) int {
	return len(people)
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

func youngest(people []Person) Person {
	youngest := 1000
	youngestPerson := Person{}
	for _, person := range people {
		if person.age < youngest {
			youngest = person.age
			youngestPerson = person
		}
	}
	return youngestPerson
}
