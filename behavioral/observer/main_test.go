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

func TestObserverPattern(t *testing.T) {
	subject := &Subject{}

	observer1 := &ConcreteObserver{name: "Papa"}
	observer2 := &ConcreteObserver{name: "Ari"}

	subject.AddObserver(observer1)
	subject.AddObserver(observer2)

	output := captureOutput(func() {
		subject.SetState("À la maison")
	})
	expectedOutput := "Papa où t'ai ? À la maison\nAri où t'ai ? À la maison\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")

	output = captureOutput(func() {
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

	output := captureOutput(func() {
		subject.SetState("Au parc")
	})
	expectedOutput := "Ari où t'ai ? Au parc\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
