package main

import (
	"encoding/csv"
	"fmt"
	"go_address_book/src/models"
	"go_address_book/src/utils"
	"os"
)

func main() {
	fmt.Println(models.Blah)
	args := os.Args
	// fmt.Println(args)

	inputFile := args[1]

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
	people := utils.ParseInput(lines)
	defer f.Close()

	// fmt.Println(people)
	find(people, cmds)
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}

func find(people []models.Person, cmds []string) {
	cmd := cmds[0]
	switch cmd {
	case "oldest":
		fmt.Println("Looking for the oldest")
		person := utils.Oldest(people)
		fmt.Println("The oldest person is ", person)
	case "youngest":
		fmt.Println("Looking for the youngest")
		person := utils.Youngest(people)
		fmt.Println("The youngest person is ", person)
	case "total":
		fmt.Println("Looking for the total number of people")
		fmt.Println("Total number of people ", utils.Total(people))
	case "city":
		fmt.Println(fmt.Sprintf("Looking for city %s", cmds[1]))
		person, err := utils.City(people, cmds[1])
		if err != nil {
			handleError(err)
		}
		fmt.Println("Looking for the person with city ", person)
	case "name":
		fmt.Println(fmt.Sprintf("Looking for name %s", cmds[1]))
		person, err := utils.Name(people, cmds[1])
		if err != nil {
			handleError(err)
		}
		fmt.Println("Looking for someone called ", person)
		// fallthrough - this is used to pass control flow to th next case
	default:
		fmt.Println("Unable to find", cmd)
	}
}
