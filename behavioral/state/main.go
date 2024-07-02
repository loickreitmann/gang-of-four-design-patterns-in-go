/*
State Example:
Allows an object to alter its behavior when its internal state
changes, appearing as if the object changed its class.
*/
package main

import (
	"fmt"
	"reflect"
)

// State Interface
type State interface {
	DoAction(context *Context)
}

// Concrete State A
type StartState struct{}

func (s *StartState) DoAction(context *Context) {
	fmt.Println("Alpha")
	context.SetState(s)
}

// Concrete State B
type StopState struct{}

func (s *StopState) DoAction(context *Context) {
	fmt.Println("Omega")
	context.SetState(s)
}

// Context
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) GetState() State {
	return c.state
}

func main() {
	context := &Context{}

	startState := &StartState{}
	startState.DoAction(context)

	fmt.Println("Current state:", reflect.TypeOf(context.GetState()))

	stopState := &StopState{}
	stopState.DoAction(context)

	fmt.Println("Current state:", reflect.TypeOf(context.GetState()))
}
