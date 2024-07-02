/*
Decorator Example:
Attaches additional responsibilities to an object
dynamically, providing a flexible alternative to
subclassing for extending functionality.
*/
package main

import "fmt"

// Component
type Notifier interface {
	Send(message string)
}

// Concrete Component
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// Decorator
type NotifierDecorator struct {
	wrapped Notifier
}

func (d *NotifierDecorator) Send(message string) {
	d.wrapped.Send(message)
}

// Concrete Decorator
type SMSNotifier struct {
	NotifierDecorator
}

func (s *SMSNotifier) Send(message string) {
	s.NotifierDecorator.Send(message)
	fmt.Println("Sending SMS:", message)
}

func main() {
	emailNotifier := &EmailNotifier{}
	smsNotifier := &SMSNotifier{NotifierDecorator{emailNotifier}}

	emailNotifier.Send("Hello Email")
	smsNotifier.Send("Hello SMS")
}
