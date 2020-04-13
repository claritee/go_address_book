package utils

import "go_address_book/src/models"

func Name(people []models.Person, name string) (models.Person, error) {
	for _, person := range people {
		if person.Name == name {
			return person, nil
		}
	}
	return models.Person{}, &models.PersonNotFoundError{}
}

func City(people []models.Person, city string) (models.Person, error) {
	for _, person := range people {
		if person.City == city {
			return person, nil
		}
	}
	return models.Person{}, &models.PersonNotFoundError{}
}

func Total(people []models.Person) int {
	return len(people)
}

func Oldest(people []models.Person) models.Person {
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

func Youngest(people []models.Person) models.Person {
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
