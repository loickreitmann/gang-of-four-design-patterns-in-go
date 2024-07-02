package main

import (
	"fmt"
	"gang_of_four_design_patterns_in_go/helper"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartStateDoAction(t *testing.T) {
	context := &Context{}
	startState := &StartState{}

	output := helper.CaptureOutput(func() {
		startState.DoAction(context)
	})

	expectedOutput := "Alpha\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
	assert.Equal(t, reflect.TypeOf(startState), reflect.TypeOf(context.GetState()), "Context state should be StartState")
}

func TestStopStateDoAction(t *testing.T) {
	context := &Context{}
	stopState := &StopState{}

	output := helper.CaptureOutput(func() {
		stopState.DoAction(context)
	})

	expectedOutput := "Omega\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
	assert.Equal(t, reflect.TypeOf(stopState), reflect.TypeOf(context.GetState()), "Context state should be StopState")
}

func TestStateTransitions(t *testing.T) {
	context := &Context{}

	startState := &StartState{}
	startState.DoAction(context)
	assert.Equal(t, reflect.TypeOf(startState), reflect.TypeOf(context.GetState()), "Context state should be StartState")

	stopState := &StopState{}
	stopState.DoAction(context)
	assert.Equal(t, reflect.TypeOf(stopState), reflect.TypeOf(context.GetState()), "Context state should be StopState")
}

func TestMainScenario(t *testing.T) {
	context := &Context{}

	output := helper.CaptureOutput(func() {
		startState := &StartState{}
		startState.DoAction(context)
		fmt.Println("Current state:", reflect.TypeOf(context.GetState()))

		stopState := &StopState{}
		stopState.DoAction(context)
		fmt.Println("Current state:", reflect.TypeOf(context.GetState()))
	})

	expectedOutput := "Alpha\nCurrent state: *main.StartState\nOmega\nCurrent state: *main.StopState\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
