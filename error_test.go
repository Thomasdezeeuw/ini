// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import (
	"errors"
	"testing"
)

func TestNewSynthaxError(t *testing.T) {
	t.Parallel()
	err := createSynthaxError(1, "error message")

	expected := "ini: synthax error on line 1: error message"
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
	regularError := errors.New("some error")
	if got := IsSynthaxError(regularError); got != false {
		t.Fatalf("Expected IsSynthaxError(%v) to return false", regularError)
	}

	synthaxError := createSynthaxError(1, "error message")
	if got := IsSynthaxError(synthaxError); got != true {
		t.Fatalf("Expected IsSynthaxError(%v) to return true", synthaxError)
	}
}

func TestCreateOverflowError(t *testing.T) {
	t.Parallel()
	err := createOverflowError("5000", "int8")
	expected := `ini: can't convert "5000" to type int8, it overflows the type`
	if err.Error() != expected {
		t.Fatalf("Expected the error message error to be %q, but got %q",
			expected, err.Error())
	}
}

func TestCreateConvertError(t *testing.T) {
	t.Parallel()
	err := createCovertionError("string", "int8")
	expected := `ini: can't convert "string" to type int8`
	if err.Error() != expected {
		t.Fatalf("Expected the error message error to be %q, but got %q",
			expected, err.Error())
	}
}
