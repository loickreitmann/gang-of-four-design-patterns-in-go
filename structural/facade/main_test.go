package main

import (
	"gang_of_four_design_patterns_in_go/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsystemAOperationA(t *testing.T) {
	subsystemA := &SubsystemA{}

	output := helper.CaptureOutput(func() {
		subsystemA.OperationA()
	})

	expectedOutput := "Subsystem A Operation\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestSubsystemBOperationB(t *testing.T) {
	subsystemB := &SubsystemB{}

	output := helper.CaptureOutput(func() {
		subsystemB.OperationB()
	})

	expectedOutput := "Subsystem B Operation\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestFacadeOperation(t *testing.T) {
	facade := &Facade{}

	output := helper.CaptureOutput(func() {
		facade.Operation()
	})

	expectedOutput := "Subsystem A Operation\nSubsystem B Operation\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
