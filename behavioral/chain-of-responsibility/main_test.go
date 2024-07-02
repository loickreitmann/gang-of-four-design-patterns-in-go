package main

import (
	"testing"

	"gang_of_four_design_patterns_in_go/helper"

	"github.com/stretchr/testify/assert"
)

func TestConcreteHandlerSetNext(t *testing.T) {
	handler1 := &ConcreteHandler{name: "one"}
	handler2 := &ConcreteHandler{name: "two"}

	handler1.SetNext(handler2)

	assert.Equal(t, handler2, handler1.next, "handler1 next should be handler2")
}

func TestConcreteHandlerHandle(t *testing.T) {
	handler1 := &ConcreteHandler{name: "one"}

	output := helper.CaptureOutput(func() {
		handler1.Handle("Process this request")
	})

	expectedOutput := "Handling request: Process this request\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestConcreteHandlerChain(t *testing.T) {
	handler1 := &ConcreteHandler{name: "one"}
	handler2 := &ConcreteHandler{name: "two"}
	handler3 := &ConcreteHandler{name: "three"}
	handler4 := &ConcreteHandler{name: "four"}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)
	handler3.SetNext(handler4)

	output := helper.CaptureOutput(func() {
		handler1.Handle("Process this request")
	})

	expectedOutput := "Handling request: Process this request\nPassing request to three\nPassing request to two\nPassing request to one\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
