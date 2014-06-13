package main

import "fmt"

// START OMIT
type Person struct {
	Name string
}

type Employee struct {
	Person // HL
	Salary int
}

func main() {
	people := []Person{}
	people = append(people, Person{"joe"})
	people = append(people, Employee{Person{"bob"}, 1000})
	for _, p := range people {
		fmt.Printf("person: %v, name: %s\n", p, p.Name)
	}
}

// END OMIT
