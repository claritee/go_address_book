package main

import (
	"database/sql"
)

type person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}

func (p *person) getPerson(db *sql.DB) error {
	return db.QueryRow("SELECT name, city, age FROM person WHERE id=$1",
		p.ID).Scan(&p.Name, &p.City, &p.Age)
}
