package cmd

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

// newTestLogger creates an in-memory logger for testing purposes.
func newTestLogger() (*slog.Logger, *bytes.Buffer) {
	var logBuffer bytes.Buffer
	var opts = slog.HandlerOptions{Level: slog.LevelDebug.Level()}
	logger := slog.New(slog.NewTextHandler(&logBuffer, &opts))
	return logger, &logBuffer
}

func TestNewHelloCmd(t *testing.T) {
	// Create a test logger and buffer to capture logs
	logger, logBuffer := newTestLogger()

	// Initialize the command
	cmd := newHelloCmd(logger)

	t.Run("Default Hello", func(t *testing.T) {
		// Prepare a buffer to capture command output
		output := &bytes.Buffer{}
		cmd.SetOut(output)
		cmd.SetArgs([]string{}) // No arguments provided

		// Execute the command
		err := cmd.Execute()
		assert.NoError(t, err)

		// Verify the output
		assert.Equal(t, "Hello World!", output.String())

		// Verify the logger captured the expected debug message
		assert.Contains(t, logBuffer.String(), "Running hello")
	})

	t.Run("Hello with Argument", func(t *testing.T) {
		// Prepare a buffer to capture command output
		output := &bytes.Buffer{}
		cmd.SetOut(output)
		cmd.SetArgs([]string{"Alice"}) // Provide a single argument

		// Execute the command
		err := cmd.Execute()
		assert.NoError(t, err)

		// Verify the output
		assert.Equal(t, "Hello Alice!", output.String())

		// Verify the logger captured the expected debug message
		assert.Contains(t, logBuffer.String(), "Running hello")
	})
}
