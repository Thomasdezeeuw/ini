// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import "fmt"

type synthaxError struct {
	lineNumber int
	msg        string
}

func (err synthaxError) Error() string {
	return fmt.Sprintf("ini: synthax error on line %d: %s",
		err.lineNumber, err.msg)
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
	return synthaxError{
		lineNumber: lineNumber,
		msg:        msg,
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
	_, ok := err.(synthaxError)
	return ok
}
