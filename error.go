package ini

import "fmt"

type synthaxError struct {
	lineNumber int
	line       string
	msg        string
}

func (s synthaxError) Error() string {
	return fmt.Sprintf("ini: synthax error on line %d. %s: %s",
		s.lineNumber, s.line, s.msg)
}

func newSynthaxError(lineNumber int, line, msg string) error {
	return synthaxError{
		lineNumber: lineNumber,
		line:       line,
		msg:        msg,
	}
}

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
