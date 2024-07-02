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

func TestConcreteHandlerSetNext(t *testing.T) {
	handler1 := &ConcreteHandler{name: "one"}
	handler2 := &ConcreteHandler{name: "two"}

	handler1.SetNext(handler2)

	assert.Equal(t, handler2, handler1.next, "handler1 next should be handler2")
}

func TestConcreteHandlerHandle(t *testing.T) {
	handler1 := &ConcreteHandler{name: "one"}

	output := captureOutput(func() {
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

	output := captureOutput(func() {
		handler1.Handle("Process this request")
	})

	expectedOutput := "Handling request: Process this request\nPassing request to three\nPassing request to two\nPassing request to one\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
