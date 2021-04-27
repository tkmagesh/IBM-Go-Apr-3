package main

import (
	"fmt"
	"math"
)

type Dimension interface {
	IHaveArea
	IHavePerimeter
}

type IHaveArea interface {
	Area() float32
}

type IHavePerimeter interface {
	Perimeter() float32
}

type Rect struct {
	length, width float32
}

func (r Rect) Area() float32 {
	return r.length * r.width
}

func (r Rect) Perimeter() float32 {
	return 2*r.length + 2*r.width
}

func (r Rect) String() string {
	return fmt.Sprintf("React {length=%v, width=%v, area=%v, perimeter=%v}", r.length, r.width, r.Area(), r.Perimeter())
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func main() {
	rect := Rect{10, 20}
	circle := Circle{15}

	/*
		printArea(rect)
		printArea(circle)

		printPerimeter(rect)
		printPerimeter(circle)
	*/

	fmt.Println(rect)
	printDimension(rect)
	printDimension(circle)
}

func printArea(o IHaveArea) {
	fmt.Println("Area = ", o.Area())
}

func printPerimeter(o IHavePerimeter) {
	fmt.Println("Perimeter = ", o.Perimeter())
}

func printDimension(d Dimension) {
	printArea(d)
	printPerimeter(d)
}

/*
	perimeter for rect = 2 * length + 2 * width
	perimter for circle = 2 * pi * radius
*/
