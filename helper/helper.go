package helper

import (
	"bytes"
	"os"
)

// Helper function to capture output from a function
func CaptureOutput(f func()) string {
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
