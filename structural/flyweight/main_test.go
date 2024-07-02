package main

import (
	"gang_of_four_design_patterns_in_go/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleDraw(t *testing.T) {
	circle := &Circle{color: "Red"}

	output := helper.CaptureOutput(func() {
		circle.Draw()
	})

	expectedOutput := "Drawing Circle with color Red\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestShapeFactoryGetCircle(t *testing.T) {
	factory := &ShapeFactory{}

	circle1 := factory.GetCircle("Red")
	circle2 := factory.GetCircle("Green")
	circle3 := factory.GetCircle("Red")

	assert.Equal(t, "Red", circle1.color, "Circle1 should have color Red")
	assert.Equal(t, "Green", circle2.color, "Circle2 should have color Green")
	assert.Equal(t, "Red", circle3.color, "Circle3 should have color Red")

	assert.Equal(t, circle1, circle3, "Circle1 and Circle3 should be the same instance")
	assert.NotEqual(t, circle1, circle2, "Circle1 and Circle2 should be different instances")
}

func TestShapeFactoryDrawCircles(t *testing.T) {
	factory := &ShapeFactory{}

	output := helper.CaptureOutput(func() {
		circle1 := factory.GetCircle("Red")
		circle1.Draw()

		circle2 := factory.GetCircle("Green")
		circle2.Draw()

		circle3 := factory.GetCircle("Red")
		circle3.Draw()
	})

	expectedOutput := "Drawing Circle with color Red\nDrawing Circle with color Green\nDrawing Circle with color Red\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
