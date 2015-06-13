// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import "fmt"

type SynthaxError struct {
	LineNumber int
	Message    string
}

func (err SynthaxError) Error() string {
	return fmt.Sprintf("ini: synthax error on line %d: %s",
		err.LineNumber, err.Message)
}

type overflowError struct {
	value string
	t     string
}

func (err overflowError) Error() string {
	return fmt.Sprintf("ini: can't convert %q to type %s, it overflows the type",
		err.value, err.t)
}

type covertionError struct {
	value string
	t     string
}

func (err covertionError) Error() string {
	return fmt.Sprintf("ini: can't convert %q to type %s", err.value, err.t)
}

func createSynthaxError(lineNumber int, msg string) error {
	return SynthaxError{
		LineNumber: lineNumber,
		Message:    msg,
	}
}

func createOverflowError(value, t string) error {
	return overflowError{
		value: value,
		t:     t,
	}
}

func createCovertionError(value, t string) error {
	return covertionError{
		value: value,
		t:     t,
	}
}

// IsSynthaxError checks if an error is a synthax error.
func IsSynthaxError(err error) bool {
	_, ok := err.(SynthaxError)
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
