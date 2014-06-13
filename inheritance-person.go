package main

import "fmt"

type Person struct {
	Name string
}

type Employee struct {
	Person
	Salary int
}

func main() {
	bob := Employee{Person{"bob"}, 1000}
	joe := Person{"joe"}
	fmt.Printf("bob's name: %s\n", bob.Name)
	fmt.Printf("joe's name: %s\n", joe.Name)
}
