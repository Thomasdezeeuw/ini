package ini

import (
	"errors"
	"testing"
)

func TestNewSynthaxError(t *testing.T) {
	t.Parallel()
	err := newSynthaxError(1, "the line", "error message")

	expected := "ini: synthax error on line 1. the line: error message"
	if err.Error() != expected {
		t.Fatalf("Expected the synthax error to be %q, but got %q",
			expected, err.Error())
	}

	if !IsSynthaxError(err) {
		t.Fatal("Expected the synthax error to be a synthax error, but it isn't")
	}
}

func TestIsSynthaxError(t *testing.T) {
	t.Parallel()
	if got := IsSynthaxError(errors.New("some error")); got != false {
		t.Fatalf("Expected synthax error to return false on regular errors")
	}

	got := IsSynthaxError(newSynthaxError(1, "the line", "error message"))
	if got != true {
		t.Fatalf("Expected synthax error to return true on synthax errors")
	}
}

func TestCreateOverflowError(t *testing.T) {
	t.Parallel()
	err := createOverflowError("5000", "int8")
	expected := `can't convert "5000" to type int8, it overflows type int8`
	if err.Error() != expected {
		t.Fatalf("Expected the error message error to be %q, but got %q",
			expected, err.Error())
	}
}

func TestCreateConvertError(t *testing.T) {
	t.Parallel()
	err := createConvertError("string", "int8")
	expected := `can't convert "string" to type int8`
	if err.Error() != expected {
		t.Fatalf("Expected the error message error to be %q, but got %q",
			expected, err.Error())
	}
}
