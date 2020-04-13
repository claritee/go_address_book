package main

import (
	"encoding/csv"
	"fmt"
	"go_address_book/models"
	"os"
	"strconv"
)

const defaultFile = "people.csv"

type PersonNotFoundError struct{}

func (e *PersonNotFoundError) Error() string {
	return "Person not found"
}

func main() {
	fmt.Println(models.Blah)
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

func parseInput(lines [][]string) []models.Person {
	people := []models.Person{}
	for _, line := range lines {
		fmt.Println(line)
		a, _ := strconv.Atoi(line[1])
		p := models.Person{
			Name: line[0],
			Age:  a,
			City: line[2],
		}
		people = append(people, p)
	}
	return people
}

func find(people []models.Person, cmds []string) {
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

func name(people []models.Person, name string) (models.Person, error) {
	for _, person := range people {
		if person.Name == name {
			return person, nil
		}
	}
	return models.Person{}, &PersonNotFoundError{}
}

func city(people []models.Person, city string) (models.Person, error) {
	for _, person := range people {
		if person.City == city {
			return person, nil
		}
	}
	return models.Person{}, &PersonNotFoundError{}
}

func total(people []models.Person) int {
	return len(people)
}

func oldest(people []models.Person) models.Person {
	oldest := 0
	oldestPerson := models.Person{}
	for _, person := range people {
		if person.Age > oldest {
			oldest = person.Age
			oldestPerson = person
		}
	}
	return oldestPerson
}

func youngest(people []models.Person) models.Person {
	youngest := 1000
	youngestPerson := models.Person{}
	for _, person := range people {
		if person.Age < youngest {
			youngest = person.Age
			youngestPerson = person
		}
	}
	return youngestPerson
}
