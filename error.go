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

type OverflowError struct {
	Value string
	Type  string
}

func (err OverflowError) Error() string {
	return fmt.Sprintf("ini: can't convert %q to type %s, it overflows the type",
		err.Value, err.Type)
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
	return OverflowError{
		Value: value,
		Type:  t,
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
	_, ok := err.(OverflowError)
	return ok
}

// IsCovertionError checks if an error is a covertion error.
func IsCovertionError(err error) bool {
	_, ok := err.(covertionError)
	return ok
}
