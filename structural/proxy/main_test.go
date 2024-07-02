package main

import (
	"gang_of_four_design_patterns_in_go/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRealSubjectRequest(t *testing.T) {
	realSubject := &RealSubject{}

	output := helper.CaptureOutput(func() {
		realSubject.Request()
	})

	expectedOutput := "RealSubject Request\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestProxyRequest(t *testing.T) {
	proxy := &Proxy{}

	output := helper.CaptureOutput(func() {
		proxy.Request()
	})

	expectedOutput := "Proxy Request\nRealSubject Request\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestProxyRequestCalledTwice(t *testing.T) {
	proxy := &Proxy{}

	output := helper.CaptureOutput(func() {
		proxy.Request()
		proxy.Request()
	})

	expectedOutput := "Proxy Request\nRealSubject Request\nProxy Request\nRealSubject Request\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
