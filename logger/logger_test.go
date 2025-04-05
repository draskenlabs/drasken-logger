package logger_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/draskenlabs/drasken-logger/logger"
	// Replace with your actual module path
)

// captureOutput temporarily redirects os.Stdout and captures anything written.
func captureOutput(f func()) string {
	// Create a pipe
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	outC := make(chan string)

	// Read output in a goroutine
	go func() {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)
		outC <- buf.String()
	}()

	// Run the function
	f()

	// Close the writer to signal EOF to reader
	_ = w.Close()
	os.Stdout = stdout // Restore original stdout

	// Return the captured output
	return <-outC
}

// TestLogger_LevelFiltering ensures the logger only logs messages >= MinLevel.
func TestLogger_LevelFiltering(t *testing.T) {
	l := logger.New(logger.WARN, false)

	output := captureOutput(func() {
		l.Info("This should not be logged") // Below WARN
		l.Warn("This is a warning")         // Included
		l.Error("This is an error")         // Included
	})

	if strings.Contains(output, "This should not be logged") {
		t.Error("Expected INFO message to be filtered out")
	}
	if !strings.Contains(output, "This is a warning") {
		t.Error("Expected WARN message to be present")
	}
	if !strings.Contains(output, "This is an error") {
		t.Error("Expected ERROR message to be present")
	}
}

// TestLogger_ColorOutput verifies that ANSI color codes are present when enabled.
func TestLogger_ColorOutput(t *testing.T) {
	l := logger.New(logger.DEBUG, true)
	l.ColorTarget = "message"

	output := captureOutput(func() {
		l.Debug("Colored debug message")
	})

	if !strings.Contains(output, "\033[36m") {
		t.Error("Expected ANSI color code in output")
	}
	if !strings.Contains(output, "Colored debug message") {
		t.Error("Expected debug message content")
	}
}

// ExampleLogger demonstrates basic usage of the logger.
func ExampleLogger() {
	l := logger.New(logger.INFO, true)
	l.Info("Starting the system")
	l.Warn("Low disk space")
	l.Error("Failed to load config")

	// This will be displayed in `pkg.go.dev`'s "Examples" tab.
	fmt.Println("Logger example ran")
}
