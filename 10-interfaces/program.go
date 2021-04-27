package main

import (
	"fmt"
	"math"
)

type IHaveArea interface {
	Area() float32
}

type Rect struct {
	length, width float32
}

func (r Rect) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func main() {
	rect := Rect{10, 20}
	circle := Circle{15}

	printArea(rect)
	printArea(circle)
}

func printArea(o IHaveArea) {
	fmt.Println("Area = ", o.Area())
}
