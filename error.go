// Copyright (C) 2015-2016 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package ini

import "fmt"

type synthaxError struct {
	LineNumber int
	Message    string
}

func (err synthaxError) Error() string {
	return fmt.Sprintf("ini: synthax error on line %d: %s",
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

func createSynthaxError(lineNumber int, msg string) error {
	return synthaxError{
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

// IsSynthaxError checks if an error is a synthax error.
func IsSynthaxError(err error) bool {
	_, ok := err.(synthaxError)
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
