package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	expected := "Specific request"
	actual := adapter.Request()

	assert.Equal(t, expected, actual, "The adapter should return the same output as the adaptee's SpecificRequest method")
}
