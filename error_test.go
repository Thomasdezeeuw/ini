package ini

import (
	"errors"
	"testing"
)

func TestNewSynthaxError(t *testing.T) {
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
	if got := IsSynthaxError(errors.New("some error")); got != false {
		t.Fatalf("Expected synthax error to return false on regular errors")
	}

	got := IsSynthaxError(newSynthaxError(1, "the line", "error message"))
	if got != true {
		t.Fatalf("Expected synthax error to return true on synthax errors")
	}
}
