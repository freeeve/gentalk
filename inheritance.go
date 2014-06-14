package main

import "fmt"

// START OMIT
type Namer interface { // HL
	Name() string // HL
} // HL

type Person struct {
	NameStr string
}

func (p Person) Name() string {
	return p.NameStr
}

type Employee struct {
	Person
	Salary int
}
// END OMIT

func main() {
	people := []Namer{}
	people = append(people, Person{"joe"})
	people = append(people, Employee{Person{"bob"}, 1000})
	for _, p := range people {
		fmt.Printf("person: %v, name: %s\n", p, p.Name())
	}
}

