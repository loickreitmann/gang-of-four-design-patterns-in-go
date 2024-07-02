package main

import (
	"bytes"
	"os"
	"strings"
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

func TestPrintCommandExecute(t *testing.T) {
	command := &PrintCommand{message: "Test Message"}

	output := captureOutput(func() {
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

	output := captureOutput(func() {
		invoker.ExecuteCommands()
	})

	expectedOutput := strings.Join(msgs, "\n") + "\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestInvokerExecuteNoCommands(t *testing.T) {
	invoker := &Invoker{}

	output := captureOutput(func() {
		invoker.ExecuteCommands()
	})

	expectedOutput := ""
	assert.Equal(t, expectedOutput, output, "Output should be empty when no commands are stored")
}
