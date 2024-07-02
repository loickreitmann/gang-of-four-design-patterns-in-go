package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock DrawingAPI for testing
type MockDrawingAPI struct {
	called       bool
	x, y, radius int
}

func (m *MockDrawingAPI) DrawCircle(x, y, radius int) {
	m.called = true
	m.x, m.y, m.radius = x, y, radius
}

func TestCircleDraw(t *testing.T) {
	mockAPI := &MockDrawingAPI{}
	circle := &Circle{x: 1, y: 2, radius: 3, drawingAPI: mockAPI}

	circle.Draw()

	assert.True(t, mockAPI.called, "DrawCircle should be called")
	assert.Equal(t, 1, mockAPI.x, "X coordinate should be 1")
	assert.Equal(t, 2, mockAPI.y, "Y coordinate should be 2")
	assert.Equal(t, 3, mockAPI.radius, "Radius should be 3")
}
