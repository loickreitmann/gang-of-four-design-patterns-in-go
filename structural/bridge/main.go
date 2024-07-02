/*
Bridge Example:
Separates an object's interface from its implementation
so the two can vary independently.
*/
package main

import "fmt"

// Implementor
type DrawingAPI interface {
	DrawCircle(x, y, radius int)
}

// Concrete Implementor A
type DrawingAPI1 struct{}

func (d *DrawingAPI1) DrawCircle(x, y, radius int) {
	fmt.Printf("API1.circle at %d:%d radius %d\n", x, y, radius)
}

// Concrete Implementor B
type DrawingAPI2 struct{}

func (d *DrawingAPI2) DrawCircle(x, y, radius int) {
	fmt.Printf("API2.circle at %d:%d radius %d\n", x, y, radius)
}

// Abstraction
type Shape interface {
	Draw()
}

// Refined Abstraction
type Circle struct {
	x, y, radius int
	drawingAPI   DrawingAPI
}

func (c *Circle) Draw() {
	c.drawingAPI.DrawCircle(c.x, c.y, c.radius)
}

func main() {
	circle1 := &Circle{1, 2, 3, &DrawingAPI1{}}
	circle2 := &Circle{5, 7, 11, &DrawingAPI2{}}

	circle1.Draw()
	circle2.Draw()
}
