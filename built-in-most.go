package main

import "fmt"

// START OMIT
type SomeEnum uint8

const (
	Foo SomeEnum = iota
	Bar
)

func main() {
	enums := []SomeEnum{}
	enums = append(enums, Foo)
	enums = append(enums, Bar)
	enums = append(enums, 7) // HL
	fmt.Println(enums)
}

// END OMIT
