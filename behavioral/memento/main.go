/*
Memento Example:
Captures and externalizes an object's internal state without
violating encapsulation, so the object can be restored to this
state later.
*/
package main

import "fmt"

// Memento
type Memento struct {
	state string
}

// Originator
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) SaveState() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) RestoreState(memento *Memento) {
	o.state = memento.state
}

func (o *Originator) GetState() string {
	return o.state
}

// Caretaker
type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) AddMemento(memento *Memento) {
	c.mementos = append(c.mementos, memento)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	return c.mementos[index]
}

func main() {
	originator := &Originator{}
	caretaker := &Caretaker{}

	originator.SetState("Shock")
	caretaker.AddMemento(originator.SaveState())

	originator.SetState("Denial")
	caretaker.AddMemento(originator.SaveState())

	originator.SetState("Anger")
	caretaker.AddMemento(originator.SaveState())

	originator.SetState("Barganing")
	caretaker.AddMemento(originator.SaveState())

	originator.SetState("Depression")
	fmt.Println("Current State:", originator.GetState())

	originator.RestoreState(caretaker.GetMemento(0))
	fmt.Println("Restored to State:", originator.GetState())

	originator.RestoreState(caretaker.GetMemento(1))
	fmt.Println("Restored to State:", originator.GetState())

	originator.RestoreState(caretaker.GetMemento(2))
	fmt.Println("Restored to State:", originator.GetState())

	originator.RestoreState(caretaker.GetMemento(3))
	fmt.Println("Restored to State:", originator.GetState())
}
