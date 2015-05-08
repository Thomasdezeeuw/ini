// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

// todo: test io stuff: http://localhost:6060/pkg/testing/iotest/.

import (
	"reflect"
	"strings"
	"testing"
	"testing/iotest"
)

func TestSectionLine(t *testing.T) {
	tests := []struct {
		line    string
		section string
	}{
		{"[section]", "section"},
		{"[section];comment", "section"},
		{"[section] ; comment", "section"},
		{"[sec;tion]", "sec;tion"},
		{"[ s e c t i o n ]", "s e c t i o n"},
	}

	for _, test := range tests {
		section, err := parseSection([]byte(test.line))
		if err != nil {
			t.Fatalf("Didn't expect parseSection(%s) to return error: '%s'",
				test.line, err.Error())
		}

		if section != test.section {
			t.Fatalf("Expected parseSection(%s) to return section: %q, but got %q",
				test.line, test.section, section)
		}
	}
}

func TestSectionLineError(t *testing.T) {
	tests := []struct {
		line   string
		errMsg string
	}{
		{"section]", "section should start with \"[\""},
		{"[section] something", "unexpected \"s\" after section closed"},
		{"[section", "unclosed section"},
	}

	for _, test := range tests {
		_, err := parseSection([]byte(test.line))
		if err == nil {
			t.Fatalf("Expected parseSection(%s) to return an error, but didn't get one",
				test.line)
		}

		if err.Error() != test.errMsg {
			t.Fatalf("Expected parseSection(%s) to return error: %q, but got %q",
				test.line, test.errMsg, err.Error())
		}
	}
}

func TestKeyValueLine(t *testing.T) {
	tests := []struct {
		line  string
		key   string
		value string
	}{
		{"key=value", "key", "value"}, // Simple.
		{"k e y=v a l u e", "k e y", "v a l u e"},
		{"key = value", "key", "value"},
		{"key=", "key", ""},
		{"key=value; comment", "key", "value"}, // Simple with comment.
		{"key=value ; comment", "key", "value"},
		{"key = value; comment", "key", "value"},
		{"key = value ; comment", "key", "value"},
		{"key=; comment", "key", ""}, // Simple only comment.
		{"key= ; comment", "key", ""},
		{"key=;", "key", ""}, // Simple empty comment.
		{"key= ;", "key", ""},
		{"key = ;", "key", ""},
		{"key =  ;", "key", ""},

		{`"key"=value`, "key", "value"}, // Double qoute.
		{`"key" = value`, "key", "value"},
		{`key="value"`, "key", "value"},
		{`key = "value"`, "key", "value"},
		{`"key"="value"`, "key", "value"},
		{`"key" = "value"`, "key", "value"},
		{`"key"=value; comment`, "key", "value"}, // Double qoute with comment.
		{`"key"=value ; comment`, "key", "value"},
		{`"key" = value; comment`, "key", "value"},
		{`"key" = value ; comment`, "key", "value"},
		{`key="value"; comment`, "key", "value"},
		{`key="value" ; comment`, "key", "value"},
		{`key = "value"; comment`, "key", "value"},
		{`key = "value" ; comment`, "key", "value"},
		{`"key"="value"; comment`, "key", "value"},
		{`"key"="value" ; comment`, "key", "value"},
		{`"key" = "value"; comment`, "key", "value"},
		{`"key" = "value" ; comment`, "key", "value"},
		{`"key"=; comment`, "key", ""}, // Double qoute only comment.
		{`"key"= ; comment`, "key", ""},
		{`"key" = ; comment`, "key", ""},
		{`key=""; comment`, "key", ""},
		{`key="" ; comment`, "key", ""},
		{`key = ""; comment`, "key", ""},
		{`key = "" ; comment`, "key", ""},
		{`"key"=""; comment`, "key", ""},
		{`"key"="" ; comment`, "key", ""},
		{`"key" = ""; comment`, "key", ""},
		{`"key" = "" ; comment`, "key", ""},
		{`"key"=;`, "key", ""}, // Double quote empty comment.
		{`"key"= ;`, "key", ""},
		{`"key" = ;`, "key", ""},
		{`key="";`, "key", ""},
		{`key="" ;`, "key", ""},
		{`key = "";`, "key", ""},
		{`key = "" ;`, "key", ""},
		{`"key"="";`, "key", ""},
		{`"key"="" ;`, "key", ""},
		{`"key" = "";`, "key", ""},
		{`"key" = "" ;`, "key", ""},

		{"'key'=value", "key", "value"}, // Single qoute.
		{"'key' = value", "key", "value"},
		{"key='value'", "key", "value"},
		{"key = 'value'", "key", "value"},
		{"'key'='value'", "key", "value"},
		{"'key' = 'value'", "key", "value"},
		{"'key'=value; comment", "key", "value"}, // Single qoute with comment.
		{"'key'=value ; comment", "key", "value"},
		{"'key' = value; comment", "key", "value"},
		{"'key' = value ; comment", "key", "value"},
		{"key='value'; comment", "key", "value"},
		{"key='value' ; comment", "key", "value"},
		{"key = 'value'; comment", "key", "value"},
		{"key = 'value' ; comment", "key", "value"},
		{"'key'='value'; comment", "key", "value"},
		{"'key'='value' ; comment", "key", "value"},
		{"'key' = 'value'; comment", "key", "value"},
		{"'key' = 'value' ; comment", "key", "value"},
		{"'key'=; comment", "key", ""}, // Single qoute only comment.
		{"'key'= ; comment", "key", ""},
		{"'key' = ; comment", "key", ""},
		{"key=''; comment", "key", ""},
		{"key='' ; comment", "key", ""},
		{"key = ''; comment", "key", ""},
		{"key = '' ; comment", "key", ""},
		{"'key'=''; comment", "key", ""},
		{"'key'='' ; comment", "key", ""},
		{"'key' = ''; comment", "key", ""},
		{"'key' = '' ; comment", "key", ""},
		{"'key'=;", "key", ""}, // Single quote empty comment.
		{"'key'= ;", "key", ""},
		{"'key' = ;", "key", ""},
		{"key='';", "key", ""},
		{"key='' ;", "key", ""},
		{"key = '';", "key", ""},
		{"key = '' ;", "key", ""},
		{"'key'='';", "key", ""},
		{"'key'='' ;", "key", ""},
		{"'key' = '';", "key", ""},
		{"'key' = '' ;", "key", ""},

		{`"=key"=value`, "=key", "value"}, // Escaped qoutes.
		{`"k\"ey"=value`, `k"ey`, "value"},
		{`key="val\"ue="`, "key", `val"ue=`},

		{"ke;y=value", "ke;y", "value"}, // Misc.
		{`k\\ey=val\\ue`, `k\ey`, `val\ue`},
		{`k\"ey=val\"ue`, `k"ey`, `val"ue`},
		{`key=val\"ue\"`, `key`, `val"ue"`},
		{`key="val\"ue\""`, `key`, `val"ue"`},
		{`\\key=value`, `\key`, "value"},
		{`"ke;y"=value`, "ke;y", "value"},
		{`key="val;ue"`, "key", "val;ue"},
		{`"ke;y"="val;ue"`, "ke;y", "val;ue"},
		{`key==value`, "key", `=value`},
		{`key=value=`, "key", `value=`},
	}

	for _, test := range tests {
		key, value, err := parseKeyValue([]byte(test.line))
		if err != nil {
			t.Fatalf("Didn't expect parseKeyValue(%s) to return error: '%s'",
				test.line, err.Error())
		}

		if key != test.key {
			t.Fatalf("Expected parseKeyValue(%s) to return key: %q, but got %q",
				test.line, test.key, key)
		}

		if value != test.value {
			t.Fatalf("Expected parseKeyValue(%s) to return value: %q, but got %q",
				test.line, test.key, value)
		}
	}
}

func TestKeyValueLineErrors(t *testing.T) {
	tests := []struct {
		line   string
		errMsg string
	}{
		{`"key'=value`, "qoute not closed"},
		{`"key=value`, "qoute not closed"},
		{`'key"=value`, "qoute not closed"},
		{`'key=value`, "qoute not closed"},
		{`key="value'`, "qoute not closed"},
		{`key="value`, "qoute not closed"},
		{`key='value"`, "qoute not closed"},
		{`key='value`, "qoute not closed"},

		{"key", "no separator found"},
		{"key value", "no separator found"},
		{`"key"`, "no separator found"},

		{`"key"value`, `unexpected "v", expected the seperator "="`},
		{`"key"val=ue`, `unexpected "v", expected the seperator "="`},
		{`"key" "2" = value`, `unexpected "\"", expected the seperator "="`},

		{"=value", "key can't be empty"},
	}

	for _, test := range tests {
		_, _, err := parseKeyValue([]byte(test.line))
		if err == nil {
			t.Fatalf("Expected parseKeyValue(%s) to return an error, but didn't get one",
				test.line)
		}

		if err.Error() != test.errMsg {
			t.Fatalf("Expected parseKeyValue(%s) to return error: %q, but got %q",
				test.line, test.errMsg, err.Error())
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		content string
		config  Config
	}{
		{"[section]", Config{Global: {}, "section": {}}},
		{"key=value", Config{Global: {"key": "value"}}},
		{"key=value; comment", Config{Global: {"key": "value"}}},
		{"; comment", Config{Global: {}}},
		{"key=value\n; comment", Config{Global: {"key": "value"}}},
		{"[section]\n\nkey=value", Config{Global: {}, "section": {"key": "value"}}},
		{"key=value\n[section]\nkey=value", Config{Global: {"key": "value"},
			"section": {"key": "value"}}},
		{"key=value\n[section]\nkey=value\n\n[section2]\nkey2 = value2",
			Config{Global: {"key": "value"}, "section": {"key": "value"},
				"section2": {"key2": "value2"}}},
	}

	for _, test := range tests {
		config, err := Parse(strings.NewReader(test.content))
		if err != nil {
			t.Fatalf("Unexpected error from Parse(%s): %s", test.content, err.Error())
		}

		if !reflect.DeepEqual(config, test.config) {
			t.Fatalf("Expected Parse(%s) to return %q, but got %q",
				test.content, config, test.config)
		}
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		content string
		errMsg  string
	}{
		{"key=value\nkey=value2", `ini: synthax error on line 2. key=value2: ` +
			`key "key" already used in section "SUPERGLOBAL"`},
		{"=value", `ini: synthax error on line 1. =value: key can't be empty`},
		{"[section", `ini: synthax error on line 1. [section: unclosed section`},
	}

	for _, test := range tests {
		_, err := Parse(strings.NewReader(test.content))
		if err == nil {
			t.Fatalf("Expected Parse(%s) to return an error, but didn't get one",
				test.content)
		}

		if err.Error() != test.errMsg {
			t.Fatalf("Expected Parse(%s) to return error: %q, but got %q",
				test.content, test.errMsg, err.Error())
		}
	}
}

func TestParseIOError(t *testing.T) {
	r := iotest.TimeoutReader(strings.NewReader("key=value\nkey2=value2"))

	_, err := Parse(r)
	if err == nil {
		t.Fatalf("Expected Parse() to return an error, but didn't get one")
	}

	errMsg := "ini: error reading: " + iotest.ErrTimeout.Error()
	if err.Error() != errMsg {
		t.Fatalf("Expected Parse() to return error: %q, but got %q",
			errMsg, err.Error())
	}
}
