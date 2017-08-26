// Copyright (C) 2015-2016 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package ini

import (
	"errors"
	"testing"
)

func TestSyntaxError(t *testing.T) {
	t.Parallel()
	regularError := errors.New("some error")
	if got := IsSyntaxError(regularError); got {
		t.Fatalf("Expected IsSyntaxError(%v) to return false", regularError)
	}

	syntaxError := createSyntaxError(1, "error message")
	if got := IsSyntaxError(syntaxError); !got {
		t.Fatalf("Expected IsSyntaxError(%v) to return true", syntaxError)
	}

	expected := "ini: syntax error on line 1: error message"
	if syntaxError.Error() != expected {
		t.Fatalf("Expected the syntax error to be %q, but got %q",
			expected, syntaxError.Error())
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
