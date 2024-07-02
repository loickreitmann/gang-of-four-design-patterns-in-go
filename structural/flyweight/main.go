/*
Flyweight Example:
Uses sharing to support large numbers of fine-grained
objects efficiently.
*/
package main

import "fmt"

// Flyweight
type Shape interface {
	Draw()
}

// Concrete Flyweight
type Circle struct {
	color string
}

func (c *Circle) Draw() {
	fmt.Println("Drawing Circle with color", c.color)
}

// Flyweight Factory
type ShapeFactory struct {
	circleMap map[string]*Circle
}

func (f *ShapeFactory) GetCircle(color string) *Circle {
	if f.circleMap == nil {
		f.circleMap = make(map[string]*Circle)
	}
	if circle, ok := f.circleMap[color]; ok {
		return circle
	}
	circle := &Circle{color: color}
	f.circleMap[color] = circle
	return circle
}

func main() {
	factory := &ShapeFactory{}

	circle1 := factory.GetCircle("Red")
	circle1.Draw()

	circle2 := factory.GetCircle("Green")
	circle2.Draw()

	circle3 := factory.GetCircle("Red")
	circle3.Draw()
}
