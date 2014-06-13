package main

import "fmt"

// START OMIT
type Wutter interface {
	Wut()
}

type Foo uint8

func (f Foo) Wut() {}

const FooBar Foo = 2

func main() {
	wutters := []Wutter{}
	wutters = append(wutters, FooBar)
	wutters = append(wutters, 7) // HL
	fmt.Println(wutters)
}

// END OMIT
