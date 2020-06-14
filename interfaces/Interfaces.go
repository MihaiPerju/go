package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	c := circle{radius: 2.}
	cArea := c.area()

	fmt.Println(cArea)

}
