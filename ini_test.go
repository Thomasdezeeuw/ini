// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestPossibleQoute(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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

func TestComplete(t *testing.T) {
	t.Parallel()
	const content = `; Configuration.
	msg="Welcome \"Bob\"" ; A welcome message
	name='http server' ;)

	; Database configuration.
	[database]
	user = "bob" ; Maybe it's not specific enough.
	password = password ; Don't tell the boss.

	; HTTP configuration.
	[http]
	port=8080
	url=example.com
	`

	c, err := Parse(strings.NewReader(content))
	if err != nil {
		t.Fatalf("Unexpected error opening testdata file: %q", err.Error())
	}

	var buf bytes.Buffer
	c.WriteTo(&buf)

	c2, err := Parse(&buf)
	if !reflect.DeepEqual(c, c2) {
		t.Fatalf("Expected %v, but got %v", c, c2)
	}
}
