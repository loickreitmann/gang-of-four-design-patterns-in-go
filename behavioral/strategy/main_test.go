package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStrategyExecute(t *testing.T) {
	strategy := &AddStrategy{}
	result := strategy.Execute(3, 4)
	assert.Equal(t, 7, result, "Add strategy should return the sum of two numbers")
}

func TestMultiplyStrategyExecute(t *testing.T) {
	strategy := &MultiplyStrategy{}
	result := strategy.Execute(3, 4)
	assert.Equal(t, 12, result, "Multiply strategy should return the product of two numbers")
}

func TestDivideStrategyExecute(t *testing.T) {
	strategy := &DivideStrategy{}
	result := strategy.Execute(8, 4)
	assert.Equal(t, 2, result, "Divide strategy should return the quotient of two numbers")
}

func TestContextSetAndExecuteStrategy(t *testing.T) {
	context := &Context{}

	context.SetStrategy(&AddStrategy{})
	result := context.ExecuteStrategy(3, 4)
	assert.Equal(t, 7, result, "Context with Add strategy should return the sum of two numbers")

	context.SetStrategy(&MultiplyStrategy{})
	result = context.ExecuteStrategy(3, 4)
	assert.Equal(t, 12, result, "Context with Multiply strategy should return the product of two numbers")

	context.SetStrategy(&DivideStrategy{})
	result = context.ExecuteStrategy(8, 4)
	assert.Equal(t, 2, result, "Context with Divide strategy should return the quotient of two numbers")
}

func TestOutputStrategy(t *testing.T) {
	context := &Context{}

	output := outputStrategy(context, &AddStrategy{})
	expectedOutput := "*main.AddStrategy Strategy: 34"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for AddStrategy")

	output = outputStrategy(context, &MultiplyStrategy{})
	expectedOutput = "*main.MultiplyStrategy Strategy: 120"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for MultiplyStrategy")

	output = outputStrategy(context, &DivideStrategy{})
	expectedOutput = "*main.DivideStrategy Strategy: 7"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format for DivideStrategy")
}
