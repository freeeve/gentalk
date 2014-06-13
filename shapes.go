package main

import (
	"fmt"
	"math"
)

// START OMIT
// maybe not the best name for this, but we'll go with it
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * math.Pi * 2
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// END OMIT

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{%f, %f}", r.Width, r.Height)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{%f}", c.Radius)
}

func main() {
	shapes := []Shape{
		Circle{Radius: 2.3},
		Rectangle{Width: 2, Height: 4},
	}
	fmt.Println("Here are some Shapes:")
	for _, shape := range shapes {
		fmt.Printf("%v area: %f\n", shape, shape.Area()) // HL
	}
}
