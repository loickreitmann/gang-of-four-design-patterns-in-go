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

func TestRealSubjectRequest(t *testing.T) {
	realSubject := &RealSubject{}

	output := captureOutput(func() {
		realSubject.Request()
	})

	expectedOutput := "RealSubject Request\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestProxyRequest(t *testing.T) {
	proxy := &Proxy{}

	output := captureOutput(func() {
		proxy.Request()
	})

	expectedOutput := "Proxy Request\nRealSubject Request\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestProxyRequestCalledTwice(t *testing.T) {
	proxy := &Proxy{}

	output := captureOutput(func() {
		proxy.Request()
		proxy.Request()
	})

	expectedOutput := "Proxy Request\nRealSubject Request\nProxy Request\nRealSubject Request\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
