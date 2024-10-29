package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.1415
}

func main() {
	var s1, s2 Shape

	s1 = Rectangle{1.0, 3.0}
	s2 = Circle{5}

	fmt.Printf("Area of this rectangle: %f\n", s1.Area())
	fmt.Printf("Area of this circle: %f\n", s2.Area())
}
