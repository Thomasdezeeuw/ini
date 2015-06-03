// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import "fmt"

type synthaxError struct {
	lineNumber int
	msg        string
}

func (s synthaxError) Error() string {
	return fmt.Sprintf("ini: synthax error on line %d: %s",
		s.lineNumber, s.msg)
}

func newSynthaxError(lineNumber int, msg string) error {
	return synthaxError{
		lineNumber: lineNumber,
		msg:        msg,
	}
}

// IsSynthaxError checks if a returned error is a synthax error.
func IsSynthaxError(err error) bool {
	_, ok := err.(synthaxError)
	return ok
}

func createOverflowError(value, t string) error {
	return fmt.Errorf("can't convert %q to type %s, it overflows type %s",
		value, t, t)
}

func createConvertError(value, t string) error {
	return fmt.Errorf("can't convert %q to type %s", value, t)
}
