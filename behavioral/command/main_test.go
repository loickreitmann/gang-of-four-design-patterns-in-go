package main

import (
	"strings"
	"testing"

	"gang_of_four_design_patterns_in_go/helper"

	"github.com/stretchr/testify/assert"
)

func TestPrintCommandExecute(t *testing.T) {
	command := &PrintCommand{message: "Test Message"}

	output := helper.CaptureOutput(func() {
		command.Execute()
	})

	expectedOutput := "Test Message\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestInvokerStoreAndExecuteCommands(t *testing.T) {
	msgs := []string{"Foo", "Bar", "Bazz"}
	invoker := &Invoker{}
	invoker.StoreCommand(&PrintCommand{message: msgs[0]})
	invoker.StoreCommand(&PrintCommand{message: msgs[1]})
	invoker.StoreCommand(&PrintCommand{message: msgs[2]})

	output := helper.CaptureOutput(func() {
		invoker.ExecuteCommands()
	})

	expectedOutput := strings.Join(msgs, "\n") + "\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestInvokerExecuteNoCommands(t *testing.T) {
	invoker := &Invoker{}

	output := helper.CaptureOutput(func() {
		invoker.ExecuteCommands()
	})

	expectedOutput := ""
	assert.Equal(t, expectedOutput, output, "Output should be empty when no commands are stored")
}
