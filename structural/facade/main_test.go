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

func TestSubsystemAOperationA(t *testing.T) {
	subsystemA := &SubsystemA{}

	output := captureOutput(func() {
		subsystemA.OperationA()
	})

	expectedOutput := "Subsystem A Operation\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestSubsystemBOperationB(t *testing.T) {
	subsystemB := &SubsystemB{}

	output := captureOutput(func() {
		subsystemB.OperationB()
	})

	expectedOutput := "Subsystem B Operation\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestFacadeOperation(t *testing.T) {
	facade := &Facade{}

	output := captureOutput(func() {
		facade.Operation()
	})

	expectedOutput := "Subsystem A Operation\nSubsystem B Operation\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
