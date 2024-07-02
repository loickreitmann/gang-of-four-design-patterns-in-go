package main

import (
	"gang_of_four_design_patterns_in_go/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObserverPattern(t *testing.T) {
	subject := &Subject{}

	observer1 := &ConcreteObserver{name: "Papa"}
	observer2 := &ConcreteObserver{name: "Ari"}

	subject.AddObserver(observer1)
	subject.AddObserver(observer2)

	output := helper.CaptureOutput(func() {
		subject.SetState("À la maison")
	})
	expectedOutput := "Papa où t'ai ? À la maison\nAri où t'ai ? À la maison\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")

	output = helper.CaptureOutput(func() {
		subject.RemoveObserver(observer1)
		subject.SetState("Au parc")
	})
	expectedOutput = "Ari où t'ai ? Au parc\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestRemoveObserver(t *testing.T) {
	subject := &Subject{}

	observer1 := &ConcreteObserver{name: "Papa"}
	observer2 := &ConcreteObserver{name: "Ari"}

	subject.AddObserver(observer1)
	subject.AddObserver(observer2)

	subject.RemoveObserver(observer1)

	output := helper.CaptureOutput(func() {
		subject.SetState("Au parc")
	})
	expectedOutput := "Ari où t'ai ? Au parc\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
