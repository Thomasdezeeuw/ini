// Copyright (C) 2015-2016 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package ini

import "fmt"

type syntaxError struct {
	LineNumber int
	Message    string
}

func (err syntaxError) Error() string {
	return fmt.Sprintf("ini: syntax error on line %d: %s",
		err.LineNumber, err.Message)
}

type overflowError struct {
	Value string
	Type  string
}

func (err overflowError) Error() string {
	return fmt.Sprintf("ini: can't convert '%s' to type %s, it overflows the type",
		err.Value, err.Type)
}

type covertionError struct {
	Value string
	Type  string
}

func (err covertionError) Error() string {
	return fmt.Sprintf("ini: can't convert '%s' to type %s", err.Value, err.Type)
}

func createSyntaxError(lineNumber int, msg string) error {
	return syntaxError{
		LineNumber: lineNumber,
		Message:    msg,
	}
}

func createOverflowError(value, t string) error {
	return overflowError{
		Value: value,
		Type:  t,
	}
}

func createCovertionError(value, t string) error {
	return covertionError{
		Value: value,
		Type:  t,
	}
}

// IsSyntaxError checks if an error is a syntax error.
func IsSyntaxError(err error) bool {
	_, ok := err.(syntaxError)
	return ok
}

// IsOverflowError checks if an error is an overflow error.
func IsOverflowError(err error) bool {
	_, ok := err.(overflowError)
	return ok
}

// IsCovertionError checks if an error is a covertion error.
func IsCovertionError(err error) bool {
	_, ok := err.(covertionError)
	return ok
}
