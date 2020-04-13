package models

var Blah = "Blah"

type Person struct {
	Name string
	Age  int
	City string
}

type PersonNotFoundError struct{}

func (e *PersonNotFoundError) Error() string {
	return "Person not found"
}
