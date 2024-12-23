package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("argument must be either Circle or Rectangle")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("argument must have an area() method")
	}
}
*/

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

/*
Implement the Perimeter() method for all the shapes and Print the perimeter using PrintPerimeter() function

Circle = 2 * pi * r
Rectangle = 2 * (h + w)
Square = 4 * s
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

func PrintPerimeter(x interface{ Perimeter() float64 }) {
	fmt.Println("Perimeter :", x.Perimeter())
}

func PrintShape(x interface {
	interface{ Area() float64 }
	interface{ Perimeter() float64 }
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should interface{ Perimeter() float64 }
}

/*
func PrintShape(x interface {
	Area() float64
	Perimeter() float64
}) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should interface{ Perimeter() float64 }
}
*/

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)

	s := Square{Side: 16}
	/*
		PrintArea(s)
		PrintPerimeter(s)
	*/
	PrintShape(s)
	// PrintArea("Ex ut ea officia labore.")
}
