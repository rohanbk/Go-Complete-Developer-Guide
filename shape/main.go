package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
	getName() string
}

type triangle struct {
	base   float64
	height float64
	name   string
}

type square struct {
	sideLength float64
	name       string
}

func (t triangle) getName() string {
	return t.name
}

func (s square) getName() string {
	return s.name
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sh shape) {
	fmt.Printf("The area of the %v is %v\n", sh.getName(), sh.getArea())
}

func main() {
	fmt.Println("This is my shape program")

	t := triangle{
		base:   2,
		height: 2,
		name:   "triangle",
	}

	s := square{
		sideLength: 2,
		name:       "square",
	}

	printArea(t)
	printArea(s)

}
