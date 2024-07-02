package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gang_of_four_design_patterns_in_go/helper"
)

func TestConcreteVisitorVisitElementA(t *testing.T) {
	visitor := &ConcreteVisitor{}
	element := &ElementA{}

	output := helper.CaptureOutput(func() {
		element.Accept(visitor)
	})

	expectedOutput := "Visiting ElementA\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for VisitElementA")
}

func TestConcreteVisitorVisitElementB(t *testing.T) {
	visitor := &ConcreteVisitor{}
	element := &ElementB{}

	output := helper.CaptureOutput(func() {
		element.Accept(visitor)
	})

	expectedOutput := "Visiting ElementB\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for VisitElementB")
}

func TestVisitorPattern(t *testing.T) {
	elements := []Element{&ElementA{}, &ElementB{}}
	visitor := &ConcreteVisitor{}

	output := helper.CaptureOutput(func() {
		for _, element := range elements {
			element.Accept(visitor)
		}
	})

	expectedOutput := "Visiting ElementA\nVisiting ElementB\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for visiting all elements")
}
