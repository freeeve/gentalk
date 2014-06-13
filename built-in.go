package main

import "fmt"

// START OMIT
type Wut string

const (
	Foo Wut = "foo"
	Bar     = "bar"
	Baz     = "baz"
)

func main() {
	wuts := []Wut{}
	wuts = append(wuts, Foo)
	wuts = append(wuts, Bar)
	wuts = append(wuts, Baz)
	fmt.Println(wuts)
}

// END OMIT
