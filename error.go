package jpath

import "fmt"

const (
	// ConfigError error in configuration
	ConfigError = 1
)

type TheError struct {
	Code    int
	Message string
}

// TheError create an error object with code and message
func PathError(code int, message string) *TheError {
	return &TheError{code, message}
}

// Error implement error interface
func (e *TheError) Error() string {
	return fmt.Sprintf("Error Code: %d || Error message: %s", e.Code, e.Message)
}
