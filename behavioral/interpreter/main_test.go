package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberInterpret(t *testing.T) {
	number := &Number{value: 5}
	result := number.Interpret()
	assert.Equal(t, 5, result, "Number interpretation should return the value itself")
}

func TestAddInterpret(t *testing.T) {
	add := &Add{
		left:  &Number{value: 3},
		right: &Number{value: 4},
	}
	result := add.Interpret()
	assert.Equal(t, 7, result, "Addition interpretation should return the sum of left and right expressions")
}

func TestSubtractInterpret(t *testing.T) {
	subtract := &Subtract{
		left:  &Number{value: 10},
		right: &Number{value: 3},
	}
	result := subtract.Interpret()
	assert.Equal(t, 7, result, "Subtraction interpretation should return the difference between left and right expressions")
}

func TestParseToken(t *testing.T) {
	token := "8"
	expression := parseToken(token)
	result := expression.Interpret()
	assert.Equal(t, 8, result, "Parsed token should be interpreted correctly as a number")
}

func TestParseExpression(t *testing.T) {
	expression := "3 4 + 2 -"
	result := parseExpression(expression).Interpret()
	assert.Equal(t, 5, result, "Parsed expression should be interpreted correctly")
}

func TestParseComplexExpression(t *testing.T) {
	expression := "10 2 8 + - 3 +"
	result := parseExpression(expression).Interpret()
	assert.Equal(t, 3, result, "Parsed complex expression should be interpreted correctly")
}
