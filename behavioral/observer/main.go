/*
Observer Example:
Defines a one-to-many dependency between objects so that when
one object changes state, all its dependents are notified and
updated automatically.
*/
package main

import "fmt"

// Subject
type Subject struct {
	observers []Observer
	state     string
}

func (s *Subject) AddObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

func (s *Subject) SetState(state string) {
	s.state = state
	s.NotifyObservers()
}

// Observer
type Observer interface {
	Update(state string)
}

// Concrete Observer
type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(state string) {
	fmt.Println(o.name, "où t'ai ?", state)
}

func main() {
	subject := &Subject{}

	observer1 := &ConcreteObserver{name: "Papa"}
	observer2 := &ConcreteObserver{name: "Ari"}

	subject.AddObserver(observer1)
	subject.AddObserver(observer2)

	subject.SetState("À la maison")
	subject.RemoveObserver(observer1)
	subject.SetState("Au parc")
}
