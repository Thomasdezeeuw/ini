// Copyright (C) 2015-2016 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package ini

import (
	"errors"
	"testing"
)

func TestSynthaxError(t *testing.T) {
	t.Parallel()
	regularError := errors.New("some error")
	if got := IsSynthaxError(regularError); got {
		t.Fatalf("Expected IsSynthaxError(%v) to return false", regularError)
	}

	synthaxError := createSynthaxError(1, "error message")
	if got := IsSynthaxError(synthaxError); !got {
		t.Fatalf("Expected IsSynthaxError(%v) to return true", synthaxError)
	}

	expected := "ini: synthax error on line 1: error message"
	if synthaxError.Error() != expected {
		t.Fatalf("Expected the synthax error to be %q, but got %q",
			expected, synthaxError.Error())
	}
}

func TestOverflowError(t *testing.T) {
	t.Parallel()
	regularError := errors.New("some error")
	if got := IsOverflowError(regularError); got {
		t.Fatalf("Expected IsOverflowError(%v) to return false", regularError)
	}

	overflowError := createOverflowError("5000", "int8")
	if got := IsOverflowError(overflowError); !got {
		t.Fatalf("Expected IsOverflowError(%v) to return true", overflowError)
	}

	expected := "ini: can't convert '5000' to type int8, it overflows the type"
	if overflowError.Error() != expected {
		t.Fatalf("Expected the error message error to be %q, but got %q",
			expected, overflowError.Error())
	}
}

func TestCovertionError(t *testing.T) {
	t.Parallel()
	regularError := errors.New("some error")
	if got := IsCovertionError(regularError); got {
		t.Fatalf("Expected IsCovertionError(%v) to return false", regularError)
	}

	covertionError := createCovertionError("string", "int8")
	if got := IsCovertionError(covertionError); !got {
		t.Fatalf("Expected IsCovertionError(%v) to return true", covertionError)
	}

	expected := "ini: can't convert 'string' to type int8"
	if covertionError.Error() != expected {
		t.Fatalf("Expected the error message error to be %q, but got %q",
			expected, covertionError.Error())
	}
}
