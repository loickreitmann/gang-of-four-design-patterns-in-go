package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDrawingAPI struct {
	called       bool
	x, y, radius int
}

func (m *MockDrawingAPI) DrawCircle(x, y, radius int) {
	m.called = true
	m.x, m.y, m.radius = x, y, radius
}

func TestCircleDrawWithMockAPI(t *testing.T) {
	mockAPI := &MockDrawingAPI{}
	circle := &Circle{x: 1, y: 2, radius: 3, drawingAPI: mockAPI}

	circle.Draw()

	assert.True(t, mockAPI.called, "DrawCircle should be called")
	assert.Equal(t, 1, mockAPI.x, "X coordinate should be 1")
	assert.Equal(t, 2, mockAPI.y, "Y coordinate should be 2")
	assert.Equal(t, 3, mockAPI.radius, "Radius should be 3")
}

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

func TestCircleDrawWithDrawingAPI1(t *testing.T) {
	api := &DrawingAPI1{}
	circle := &Circle{x: 5, y: 6, radius: 7, drawingAPI: api}

	output := captureOutput(func() {
		circle.Draw()
	})

	expectedOutput := "API1.circle at 5:6 radius 7\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestCircleDrawWithDrawingAPI2(t *testing.T) {
	api := &DrawingAPI2{}
	circle := &Circle{x: 8, y: 9, radius: 10, drawingAPI: api}

	output := captureOutput(func() {
		circle.Draw()
	})

	expectedOutput := "API2.circle at 8:9 radius 10\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
