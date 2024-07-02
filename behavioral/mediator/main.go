/*
Mediator Example:
Defines an object that encapsulates how a set of objects
interact, promoting loose coupling.
*/
package main

import "fmt"

// Mediator Interface
type Mediator interface {
	Send(message string, colleague Colleague)
}

// Colleague Interface
type Colleague interface {
	SetMediator(mediator Mediator)
	Send(message string)
	Receive(message string)
}

// Concrete Colleague
type ConcreteColleague struct {
	mediator Mediator
	name     string
}

func (c *ConcreteColleague) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ConcreteColleague) Send(message string) {
	fmt.Println(c.name, "sending message:", message)
	c.mediator.Send(message, c)
}

func (c *ConcreteColleague) Receive(message string) {
	fmt.Println(c.name, "received message:", message)
}

// Concrete Mediator
type ConcreteMediator struct {
	colleague1 Colleague
	colleague2 Colleague
}

func (m *ConcreteMediator) SetColleague1(colleague Colleague) {
	m.colleague1 = colleague
}

func (m *ConcreteMediator) SetColleague2(colleague Colleague) {
	m.colleague2 = colleague
}

func (m *ConcreteMediator) Send(message string, colleague Colleague) {
	if colleague == m.colleague1 {
		m.colleague2.Receive(message)
	} else {
		m.colleague1.Receive(message)
	}
}

func main() {
	mediator := &ConcreteMediator{}

	colleague1 := &ConcreteColleague{name: "Obi Wan"}
	colleague2 := &ConcreteColleague{name: "Anakin"}

	colleague1.SetMediator(mediator)
	colleague2.SetMediator(mediator)

	mediator.SetColleague1(colleague1)
	mediator.SetColleague2(colleague2)

	colleague1.Send("Hello, there.")
	colleague2.Send("This is where the fun begins!")
}
