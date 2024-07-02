package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to capture output
func captureOutput(f func()) string {
	var buf bytes.Buffer
	writer := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = writer

	buf.ReadFrom(r)
	return buf.String()
}

func TestConcreteVisitorVisitElementA(t *testing.T) {
	visitor := &ConcreteVisitor{}
	element := &ElementA{}

	output := captureOutput(func() {
		element.Accept(visitor)
	})

	expectedOutput := "Visiting ElementA\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for VisitElementA")
}

func TestConcreteVisitorVisitElementB(t *testing.T) {
	visitor := &ConcreteVisitor{}
	element := &ElementB{}

	output := captureOutput(func() {
		element.Accept(visitor)
	})

	expectedOutput := "Visiting ElementB\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for VisitElementB")
}

func TestVisitorPattern(t *testing.T) {
	elements := []Element{&ElementA{}, &ElementB{}}
	visitor := &ConcreteVisitor{}

	output := captureOutput(func() {
		for _, element := range elements {
			element.Accept(visitor)
		}
	})

	expectedOutput := "Visiting ElementA\nVisiting ElementB\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for visiting all elements")
}
