/*
Chain of Responsibility Example:
Passes a request along a chain of handlers, allowing multiple
objects a chance to handle the request.
*/
package main

import "fmt"

// Handler
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// Concrete Handler
type ConcreteHandler struct {
	name string
	next Handler
}

func (h *ConcreteHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
		fmt.Printf("Passing request to %v\n", h.name)
	} else {
		fmt.Println("Handling request:", request)
	}
}

func main() {
	handler1 := &ConcreteHandler{name: "one"}
	handler2 := &ConcreteHandler{name: "two"}
	handler3 := &ConcreteHandler{name: "three"}
	handler4 := &ConcreteHandler{name: "four"}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)
	handler3.SetNext(handler4)

	handler1.Handle("Process this request")
}
