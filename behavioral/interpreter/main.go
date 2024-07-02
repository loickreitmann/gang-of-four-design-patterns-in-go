/*
Interpreter Example:
Defines a representation for a language's grammar and uses
an interpreter to interpret sentences in the language.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Expression
type Expression interface {
	Interpret() int
}

// Number
type Number struct {
	value int
}

func (n *Number) Interpret() int {
	return n.value
}

// Add
type Add struct {
	left, right Expression
}

func (a *Add) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

// Subtract
type Subtract struct {
	left, right Expression
}

func (s *Subtract) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}

// Parse
func parseToken(token string) Expression {
	value, _ := strconv.Atoi(token)
	return &Number{value: value}
}

// Parse
func parseExpression(expression string) Expression {
	tokens := strings.Fields(expression)
	stack := []Expression{}

	for _, token := range tokens {
		switch token {
		case "+":
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, &Add{left, right})
		case "-":
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, &Subtract{left, right})
		default:
			stack = append(stack, parseToken(token))
		}
	}

	return stack[0]
}

func main() {
	expression := "3 4 + 2 -"
	result := parseExpression(expression).Interpret()
	fmt.Println("Result:", result)
}
