/*
Proxy Example:
Provides a surrogate or placeholder for another object
to control access to it.
*/
package main

import "fmt"

// Subject
type Subject interface {
	Request()
}

// Real Subject
type RealSubject struct{}

func (r *RealSubject) Request() {
	fmt.Println("RealSubject Request")
}

// Proxy
type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Request() {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	fmt.Println("Proxy Request")
	p.realSubject.Request()
}

func main() {
	proxy := &Proxy{}
	proxy.Request()
}
