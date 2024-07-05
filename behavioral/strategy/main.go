/*
Strategy Example:
Defines a family of algorithms, encapsulates each one, and makes
them interchangeable. Strategy lets the algorithm vary independently 
from clients that use it.
*/
package main

import (
	"fmt"
	"reflect"
)

// Strategy Interface
type Strategy interface {
	Execute(a, b int) int
}

// Concrete Strategy A
type AddStrategy struct{}

func (a *AddStrategy) Execute(x, y int) int {
	return x + y
}

// Concrete Strategy B
type MultiplyStrategy struct{}

func (m *MultiplyStrategy) Execute(x, y int) int {
	return x * y
}

// Concrete Strategy C
type DivideStrategy struct{}

func (d *DivideStrategy) Execute(x, y int) int {
	return x / y
}

// Context
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func outputStrategy(context *Context, strat Strategy) string {
	context.SetStrategy(strat)
	return fmt.Sprint(reflect.TypeOf(context.strategy), " Strategy: ", context.ExecuteStrategy(30, 4))
}

func main() {
	context := &Context{}

	fmt.Println(outputStrategy(context, &AddStrategy{}))

	fmt.Println(outputStrategy(context, &MultiplyStrategy{}))

	fmt.Println(outputStrategy(context, &DivideStrategy{}))
}
