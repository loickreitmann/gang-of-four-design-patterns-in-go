/*
Adapter Example:
Converts the interface of a class into another interface
that clients expect, allowing classes to work together
that couldn't otherwise because of incompatible interfaces.
*/
package main

import "fmt"

// Target Interface
type ITarget interface {
	Request() string
}

// Adaptee
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Adapter
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	fmt.Println(adapter.Request())
}
