package jpath

import "fmt"

const (
	ConfigError       = 1
	ParserError       = 2
)

type JpathError struct {
	Code    int
	Message string
}

func PathError(code int, message string) *JpathError {
	return &JpathError{code, message}
}

func (e *JpathError) Error() string {
	return fmt.Sprintf("Error Code: %d || Error message: %s", e.Code, e.Message)
}
