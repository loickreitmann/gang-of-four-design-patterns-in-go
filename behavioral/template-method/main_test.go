package main

import (
	"gang_of_four_design_patterns_in_go/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcreteClassPrimitiveOperation1(t *testing.T) {
	class := &ConcreteClass{}

	output := helper.CaptureOutput(func() {
		class.PrimitiveOperation1()
	})

	expectedOutput := "PrimitiveOperation1\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for PrimitiveOperation1")
}

func TestConcreteClassPrimitiveOperation2(t *testing.T) {
	class := &ConcreteClass{}

	output := helper.CaptureOutput(func() {
		class.PrimitiveOperation2()
	})

	expectedOutput := "PrimitiveOperation2\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for PrimitiveOperation2")
}

func TestConcreteClassTemplateMethod(t *testing.T) {
	class := &ConcreteClass{}

	output := helper.CaptureOutput(func() {
		class.TemplateMethod()
	})

	expectedOutput := "PrimitiveOperation1\nPrimitiveOperation2\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for TemplateMethod")
}
