package utils

import (
	"fmt"
	"go_address_book/models"
	"strconv"
)

func ParseInput(lines [][]string) []models.Person {
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
