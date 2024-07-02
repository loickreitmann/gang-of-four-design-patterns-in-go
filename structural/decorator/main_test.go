package main

import (
	"gang_of_four_design_patterns_in_go/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailNotifierSend(t *testing.T) {
	emailNotifier := &EmailNotifier{}

	output := helper.CaptureOutput(func() {
		emailNotifier.Send("Hello Email")
	})

	expectedOutput := "Sending Email: Hello Email\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestSMSNotifierSend(t *testing.T) {
	emailNotifier := &EmailNotifier{}
	smsNotifier := &SMSNotifier{NotifierDecorator{emailNotifier}}

	output := helper.CaptureOutput(func() {
		smsNotifier.Send("Hello SMS")
	})

	expectedOutput := "Sending Email: Hello SMS\nSending SMS: Hello SMS\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
