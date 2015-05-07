// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetKeysAlpha(t *testing.T) {
	input := map[string]string{
		"a": "a",
		"b": "b",
		"f": "f",
		"e": "e",
		"d": "d",
		"c": "c",
	}

	got := getKeysAlpha(input)
	expects := []string{"a", "b", "c", "d", "e", "f"}

	if len(got) != len(expects) {
		t.Fatalf("Expected getKeysAlpha(%v) to return %v, got %v",
			input, expects, got)
	}

	for i, k := range got {
		if expected := expects[i]; k != expected {
			t.Fatalf("Expected getKeysAlpha(%v) to return %v, got %v",
				input, expects, got)
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

func TestLoadErrors(t *testing.T) {
	tests := []struct {
		input    string
		expected Config
		err      error
		altCheck bool
	}{
		{"testdata/malformed.cfg", nil,
			newSynthaxError(2, "error = 'oops!\"", "qoute not closed"), false},
		{"testdata", nil, errFile, false},
		{"testdata/notfound", nil, nil, true},
	}

	for _, test := range tests {
		c, err := Load(test.input)

		if test.altCheck {
			if !os.IsNotExist(err) {
				t.Fatalf("Expected Load(%q) the return error to be file doesn't exists, but got %q",
					test.input, err.Error())
			}
		} else {
			if err == nil && test.err != nil {
				t.Fatalf(`Expected Load(%q) to return error "%v", but didn't got one`,
					test.input, test.err)
			} else if err.Error() != test.err.Error() {
				t.Fatalf(`Expected Load(%q) to return error "%v", but got "%v"`,
					test.input, test.err, err)
			}
		}

		if c != nil {
			t.Fatalf("Expected Load(%q) the return config to be nil, got %v",
				test.input, c)
		}
	}
}

func TestConfigString(t *testing.T) {
	var expectedConfigString = `key1=value1` + "\n" +
		`key2"=value2` + "\n" +
		`key3=value"3` + "\n\n" +

		`[Section1]` + "\n" +
		`key1=value1` + "\n" +
		`key2"=value2` + "\n" +
		`key3=value"3` + "\n\n" +

		`[Section2]` + "\n" +
		`key1=value1` + "\n" +
		`key2"=value2` + "\n" +
		`key3=value"3` + "\n\n"

	section := map[string]string{
		"key1":  "value1",
		`key2"`: "value2",
		"key3":  `value"3`,
	}

	c := Config{
		Global:     section,
		"Section1": section,
		"Section2": section,
	}

	if got := c.String(); got != expectedConfigString {
		t.Fatalf("Expected Config{}.String() to return %q, but got %q",
			expectedConfigString, got)
	}
}

func TestComplete(t *testing.T) {
	t.Skip()
	c, err := Load("testdata/config.ini")
	if err != nil {
		t.Fatal(err)
	}

	// Create a temp file
	tmpF, err := ioutil.TempFile("", "INI-TEST")
	if err != nil {
		t.Fatal(err)
	}

	tmpPath := tmpF.Name()

	// Write our config and close the file
	tmpF.WriteString(c.String())
	tmpF.Close()

	// Reopen the tmp config
	f, err := os.Open(tmpPath)
	if err != nil {
		t.Fatal(err)
	}

	// Parse the temp config file
	c2, err := Parse(f)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	// Compare configurations
	if len(c) != len(c2) {
		t.Fatalf("Length doesn't match, got %d and %d", len(c), len(c2))
	}

	for sectionName, loadedSection := range c {
		for key, loadedValue := range loadedSection {
			parsedValue := c2[sectionName][key]
			if loadedValue != parsedValue {
				t.Fatalf("Expected, in section %q, the keys %q to be the same, but got %q and %q",
					sectionName, key, loadedValue, parsedValue)
			}
		}
	}

	// Cleanup
	if err := os.Remove(tmpPath); err != nil {
		t.Fatal(err)
	}
}
