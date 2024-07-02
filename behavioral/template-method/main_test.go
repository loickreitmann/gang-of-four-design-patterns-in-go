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

func TestConcreteClassPrimitiveOperation1(t *testing.T) {
	class := &ConcreteClass{}

	output := captureOutput(func() {
		class.PrimitiveOperation1()
	})

	expectedOutput := "PrimitiveOperation1\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for PrimitiveOperation1")
}

func TestConcreteClassPrimitiveOperation2(t *testing.T) {
	class := &ConcreteClass{}

	output := captureOutput(func() {
		class.PrimitiveOperation2()
	})

	expectedOutput := "PrimitiveOperation2\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for PrimitiveOperation2")
}

func TestConcreteClassTemplateMethod(t *testing.T) {
	class := &ConcreteClass{}

	output := captureOutput(func() {
		class.TemplateMethod()
	})

	expectedOutput := "PrimitiveOperation1\nPrimitiveOperation2\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for TemplateMethod")
}
