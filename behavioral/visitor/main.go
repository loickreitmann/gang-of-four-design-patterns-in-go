/*
Visitor Example:
Represents an operation to be performed on elements of an object
structure, allowing new operations to be defined without changing
the classes of the elements on which it operates.
*/
package main

import "fmt"

// Visitor
type Visitor interface {
	VisitElementA(element *ElementA)
	VisitElementB(element *ElementB)
}

// Element Interface
type Element interface {
	Accept(visitor Visitor)
}

// Concrete Element A
type ElementA struct{}

func (e *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(e)
}

// Concrete Element B
type ElementB struct{}

func (e *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(e)
}

// Concrete Visitor
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitElementA(element *ElementA) {
	fmt.Println("Visiting ElementA")
}

func (v *ConcreteVisitor) VisitElementB(element *ElementB) {
	fmt.Println("Visiting ElementB")
}

func main() {
	elements := []Element{&ElementA{}, &ElementB{}}
	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}
