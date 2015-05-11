// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestPossibleQoute(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`my "key"`, `"my \"key\""`},
		{`my 'key'`, `"my 'key'"`},
	}

	for _, test := range tests {
		got := possibleQoute(test.input)
		if got != test.expected {
			t.Fatalf("Expected possibleQoute(`%s`) to return `%s`, but got `%s`",
				test.input, test.expected, got)
		}
	}
}

func TestGetKeysAlpha(t *testing.T) {
	input := map[string]string{
		"a": "a",
		"b": "b",
		"f": "f",
		"e": "e",
		"d": "d",
		"c": "c",
	}

	got := getMapsKeysAlpha(input)
	expects := []string{"a", "b", "c", "d", "e", "f"}

	if len(got) != len(expects) {
		t.Fatalf("Expected getMapsKeysAlpha(%v) to return %v, got %v",
			input, expects, got)
	}

	for i, k := range got {
		if expected := expects[i]; k != expected {
			t.Fatalf("Expected getMapsKeysAlpha(%v) to return %v, got %v",
				input, expects, got)
		}
	}
}

func TestGetConfigSectionsAlpha(t *testing.T) {
	c := Config{Global: {}, "section1": {}, "section2": {}}

	got := getConfigSectionsAlpha(c)
	expects := []string{Global, "section1", "section2"}

	if len(got) != len(expects) {
		t.Fatalf("Expected getConfigSectionsAlpha(%v) to return %v, got %v",
			c, expects, got)
	}

	for i, k := range got {
		if expected := expects[i]; k != expected {
			t.Fatalf("Expected getConfigSectionsAlpha(%v) to return %v, got %v",
				c, expects, got)
		}
	}
}

func TestLoad(t *testing.T) {
	const input = "testdata/config.ini"
	c, err := Load(input)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		section  string
		key      string
		expected string
	}{
		{Global, "name", "http server"},
		{Global, "msg", "Welcome \"Bob\""},
		{"http", "port", "8080"},
		{"http", "url", "example.com"},
		{"database", "user", "bob"},
		{"database", "password", "password"},
	}

	for _, test := range tests {
		if got := c[test.section][test.key]; got != test.expected {
			t.Fatalf("Expected Load(%q) to return config[%q][%q] to be %q, but got %q",
				input, test.section, test.key, test.expected, got)
		}
	}
}

func TestComplete(t *testing.T) {
	c, err := Load("testdata/config.ini")
	if err != nil {
		t.Fatalf("Unexpected error opening testdata file: %q", err.Error())
	}

	f, err := ioutil.TempFile("", "ini")
	if err != nil {
		t.Fatalf("Unexpected error opening tempfile: %q", err.Error())
	}
	c.WriteTo(f)
	f.Close()

	tmpPath := f.Name()
	c2, err := Load(tmpPath)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	if !reflect.DeepEqual(c, c2) {
		t.Fatalf("Expected %v, but got %v", c, c2)
	}

	if err := os.Remove(tmpPath); err != nil {
		t.Fatalf("Unexpected error removing tempfile: %q", err.Error())
	}
}
