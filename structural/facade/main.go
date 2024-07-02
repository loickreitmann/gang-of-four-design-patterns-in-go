/*
Facade Example:
Provides a unified interface to a set of interfaces in
a subsystem, making the subsystem easier to use.
*/
package main

import "fmt"

// Subsystem A
type SubsystemA struct{}

func (a *SubsystemA) OperationA() {
	fmt.Println("Subsystem A Operation")
}

// Subsystem B
type SubsystemB struct{}

func (b *SubsystemB) OperationB() {
	fmt.Println("Subsystem B Operation")
}

// Facade
type Facade struct {
	a SubsystemA
	b SubsystemB
}

func (f *Facade) Operation() {
	f.a.OperationA()
	f.b.OperationB()
}

func main() {
	facade := &Facade{}
	facade.Operation()
}
