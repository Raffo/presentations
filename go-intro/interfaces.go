package main

import (
	"fmt"
	"math"
)

type Measurer interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func printArea(m Measurer) {
	fmt.Printf("Area: %v\n", m.area())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	printArea(r)
	printArea(c)
}
