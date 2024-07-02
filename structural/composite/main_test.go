package main

import (
	"gang_of_four_design_patterns_in_go/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeveloperShowDetails(t *testing.T) {
	dev := &Developer{Name: "John"}

	output := helper.CaptureOutput(func() {
		dev.ShowDetails()
	})

	expectedOutput := "Developer: John\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestManagerShowDetails(t *testing.T) {
	manager := &Manager{Name: "Alice"}

	output := helper.CaptureOutput(func() {
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

	output := helper.CaptureOutput(func() {
		dir.ShowDetails()
	})

	expectedOutput := "Developer: John\nDeveloper: Jane\nManager: Alice\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
