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

func TestDeveloperShowDetails(t *testing.T) {
	dev := &Developer{Name: "John"}

	output := captureOutput(func() {
		dev.ShowDetails()
	})

	expectedOutput := "Developer: John\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestManagerShowDetails(t *testing.T) {
	manager := &Manager{Name: "Alice"}

	output := captureOutput(func() {
		manager.ShowDetails()
	})

	expectedOutput := "Manager: Alice\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestDirectoryShowDetails(t *testing.T) {
	dev1 := &Developer{Name: "John"}
	dev2 := &Developer{Name: "Jane"}
	manager := &Manager{Name: "Alice"}

	dir := &Directory{
		Employees: []Employee{dev1, dev2, manager},
	}

	output := captureOutput(func() {
		dir.ShowDetails()
	})

	expectedOutput := "Developer: John\nDeveloper: Jane\nManager: Alice\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
