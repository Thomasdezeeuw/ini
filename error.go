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

