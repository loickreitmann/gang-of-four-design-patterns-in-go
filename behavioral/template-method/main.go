/*
Template Method Example:
Defines the skeleton of an algorithm in an operation, deferring
some steps to subclasses.
*/
package main

import "fmt"

// Abstract Class
type AbstractClass interface {
	PrimitiveOperation1()
	PrimitiveOperation2()
	TemplateMethod()
}

// Concrete Class
type ConcreteClass struct{}

func (c *ConcreteClass) PrimitiveOperation1() {
	fmt.Println("PrimitiveOperation1")
}

func (c *ConcreteClass) PrimitiveOperation2() {
	fmt.Println("PrimitiveOperation2")
}

func (c *ConcreteClass) TemplateMethod() {
	c.PrimitiveOperation1()
	c.PrimitiveOperation2()
}

func main() {
	class := &ConcreteClass{}
	class.TemplateMethod()
}
