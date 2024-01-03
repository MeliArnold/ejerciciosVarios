package main

import (
	"fmt"
	"time"
)

// Estructura Person
type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

// Estructura Employee con composición de la estructura Person
type Employee struct {
	ID       int
	Position string
	Person   Person
}

// Método PrintEmployee asociado a la estructura Employee
func (e Employee) PrintEmployee() {
	fmt.Printf("ID: %d\nPosition: %s\nPerson ID: %d\nPerson Name: %s\nDate of Birth: %s\n",
		e.ID, e.Position, e.Person.ID, e.Person.Name, e.Person.DateOfBirth.Format("2006-01-02"))
}

func main() {
	// Instanciar una Person
	person := Person{
		ID:          1,
		Name:        "John Doe",
		DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	// Instanciar un Employee con la Person previamente creada
	employee := Employee{
		ID:       1001,
		Position: "Manager",
		Person:   person,
	}

	// Ejecutar el método PrintEmployee del Employee creado
	employee.PrintEmployee()
}
