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

func TestEmailNotifierSend(t *testing.T) {
	emailNotifier := &EmailNotifier{}

	output := captureOutput(func() {
		emailNotifier.Send("Hello Email")
	})

	expectedOutput := "Sending Email: Hello Email\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestSMSNotifierSend(t *testing.T) {
	emailNotifier := &EmailNotifier{}
	smsNotifier := &SMSNotifier{NotifierDecorator{emailNotifier}}

	output := captureOutput(func() {
		smsNotifier.Send("Hello SMS")
	})

	expectedOutput := "Sending Email: Hello SMS\nSending SMS: Hello SMS\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
