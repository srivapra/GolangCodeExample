// An interface is a set of method signature. It contains only the
// signatures of the functions, but not their implementation.

// Why do we need interfaces?

// Let's take a real world example of geometrical shape (like circle, rectangle, pentagon
// hexagon, square	etc) they have common functionalities (like area).
// So we make an interface and put all these common functionalities into this interface.
// and let's geometical shape implements these functionalities in their own way.
// In this way, we don't need to change the code, only we need to implemenet
// the functionalities.

package main

import (
	"fmt"
	"math"
)

type GeometricalShape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func PrintArea(g GeometricalShape) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	c := circle{radius: 3}
	r := rectangle{width: 4, height: 5}
	PrintArea(c)
	PrintArea(r)
}
