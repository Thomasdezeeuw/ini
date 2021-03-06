// Copyright (C) 2015-2016 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package ini

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"testing/iotest"
)

type ParseTest struct {
	content string
	config  Config
}

func TestParse(t *testing.T) {
	t.Parallel()
	tests := []ParseTest{
		{"", Config{Global: {}}},
		{"; comment", Config{Global: {}}},
		{"key=value\n; comment", Config{Global: {"key": "value"}}},
		{"[section]\n\nkey=value", Config{Global: {}, "section": {"key": "value"}}},
		{"key=value\n[section]\nkey=value", Config{Global: {"key": "value"},
			"section": {"key": "value"}}},
		{"key=value\n[section]\nkey=value\n\n[section2]\nkey2 = value2",
			Config{Global: {"key": "value"}, "section": {"key": "value"},
				"section2": {"key2": "value2"}}},
	}

	if err := testParser(tests); err != nil {
		t.Fatal(err.Error())
	}
}

func TestParseKeyValueDoubleQoute(t *testing.T) {
	t.Parallel()
	tests := []ParseTest{
		{"key=value", Config{Global: {"key": "value"}}}, // Simple.
		{"k e y=v a l u e", Config{Global: {"k e y": "v a l u e"}}},
		{"key = value", Config{Global: {"key": "value"}}},
		{"key=", Config{Global: {"key": ""}}},
		{"key=value; comment", Config{Global: {"key": "value"}}}, // Simple with comment.
		{"key=value ; comment", Config{Global: {"key": "value"}}},
		{"key = value; comment", Config{Global: {"key": "value"}}},
		{"key = value ; comment", Config{Global: {"key": "value"}}},
		{"key=value# comment", Config{Global: {"key": "value"}}},
		{"key=value # comment", Config{Global: {"key": "value"}}},
		{"key = value# comment", Config{Global: {"key": "value"}}},
		{"key = value # comment", Config{Global: {"key": "value"}}},
		{"key=; comment", Config{Global: {"key": ""}}}, // Simple only comment.
		{"key= ; comment", Config{Global: {"key": ""}}},
		{"key=# comment", Config{Global: {"key": ""}}},
		{"key= # comment", Config{Global: {"key": ""}}},
		{"key=;", Config{Global: {"key": ""}}}, // Simple empty comment.
		{"key= ;", Config{Global: {"key": ""}}},
		{"key = ;", Config{Global: {"key": ""}}},
		{"key =  ;", Config{Global: {"key": ""}}},
		{`"key"=value`, Config{Global: {"key": "value"}}}, // Double qoute.
		{`"key" = value`, Config{Global: {"key": "value"}}},
		{`key="value"`, Config{Global: {"key": "value"}}},
		{`key = "value"`, Config{Global: {"key": "value"}}},
		{`"key"="value"`, Config{Global: {"key": "value"}}},
		{`"key" = "value"`, Config{Global: {"key": "value"}}},
		{`"key"=value; comment`, Config{Global: {"key": "value"}}}, // Double qoute with comment.
		{`"key"=value ; comment`, Config{Global: {"key": "value"}}},
		{`"key" = value; comment`, Config{Global: {"key": "value"}}},
		{`"key" = value ; comment`, Config{Global: {"key": "value"}}},
		{`key="value"; comment`, Config{Global: {"key": "value"}}},
		{`key="value" ; comment`, Config{Global: {"key": "value"}}},
		{`key = "value"; comment`, Config{Global: {"key": "value"}}},
		{`key = "value" ; comment`, Config{Global: {"key": "value"}}},
		{`"key"="value"; comment`, Config{Global: {"key": "value"}}},
		{`"key"="value" ; comment`, Config{Global: {"key": "value"}}},
		{`"key" = "value"; comment`, Config{Global: {"key": "value"}}},
		{`"key" = "value" ; comment`, Config{Global: {"key": "value"}}},
		{`"key"=value# comment`, Config{Global: {"key": "value"}}},
		{`"key"=value # comment`, Config{Global: {"key": "value"}}},
		{`"key" = value# comment`, Config{Global: {"key": "value"}}},
		{`"key" = value # comment`, Config{Global: {"key": "value"}}},
		{`key="value"# comment`, Config{Global: {"key": "value"}}},
		{`key="value" # comment`, Config{Global: {"key": "value"}}},
		{`key = "value"# comment`, Config{Global: {"key": "value"}}},
		{`key = "value" # comment`, Config{Global: {"key": "value"}}},
		{`"key"="value"# comment`, Config{Global: {"key": "value"}}},
		{`"key"="value" # comment`, Config{Global: {"key": "value"}}},
		{`"key" = "value"# comment`, Config{Global: {"key": "value"}}},
		{`"key" = "value" # comment`, Config{Global: {"key": "value"}}},
		{`"key"=; comment`, Config{Global: {"key": ""}}}, // Double qoute only comment.
		{`"key"= ; comment`, Config{Global: {"key": ""}}},
		{`"key" = ; comment`, Config{Global: {"key": ""}}},
		{`key=""; comment`, Config{Global: {"key": ""}}},
		{`key="" ; comment`, Config{Global: {"key": ""}}},
		{`key = ""; comment`, Config{Global: {"key": ""}}},
		{`key = "" ; comment`, Config{Global: {"key": ""}}},
		{`"key"=""; comment`, Config{Global: {"key": ""}}},
		{`"key"="" ; comment`, Config{Global: {"key": ""}}},
		{`"key" = ""; comment`, Config{Global: {"key": ""}}},
		{`"key" = "" ; comment`, Config{Global: {"key": ""}}},
		{`"key"=# comment`, Config{Global: {"key": ""}}},
		{`"key"= # comment`, Config{Global: {"key": ""}}},
		{`"key" = # comment`, Config{Global: {"key": ""}}},
		{`key=""# comment`, Config{Global: {"key": ""}}},
		{`key="" # comment`, Config{Global: {"key": ""}}},
		{`key = ""# comment`, Config{Global: {"key": ""}}},
		{`key = "" # comment`, Config{Global: {"key": ""}}},
		{`"key"=""# comment`, Config{Global: {"key": ""}}},
		{`"key"="" # comment`, Config{Global: {"key": ""}}},
		{`"key" = ""# comment`, Config{Global: {"key": ""}}},
		{`"key" = "" # comment`, Config{Global: {"key": ""}}},
		{`"key"=;`, Config{Global: {"key": ""}}}, // Double quote empty comment.
		{`"key"= ;`, Config{Global: {"key": ""}}},
		{`"key" = ;`, Config{Global: {"key": ""}}},
		{`key="";`, Config{Global: {"key": ""}}},
		{`key="" ;`, Config{Global: {"key": ""}}},
		{`key = "";`, Config{Global: {"key": ""}}},
		{`key = "" ;`, Config{Global: {"key": ""}}},
		{`"key"="";`, Config{Global: {"key": ""}}},
		{`"key"="" ;`, Config{Global: {"key": ""}}},
		{`"key" = "";`, Config{Global: {"key": ""}}},
		{`"key" = "" ;`, Config{Global: {"key": ""}}},
		{`" key " = " value "`, Config{Global: {" key ": " value "}}}, // Keep whitespace.
		{`"	key	" = "	value	"`, Config{Global: {"	key	": "	value	"}}},
	}

	if err := testParser(tests); err != nil {
		t.Fatal(err.Error())
	}
}

func TestParseKeyValueSingleQoute(t *testing.T) {
	t.Parallel()
	tests := []ParseTest{
		{"'key'=value", Config{Global: {"key": "value"}}}, // Single qoute.
		{"'key' = value", Config{Global: {"key": "value"}}},
		{"key='value'", Config{Global: {"key": "value"}}},
		{"key = 'value'", Config{Global: {"key": "value"}}},
		{"'key'='value'", Config{Global: {"key": "value"}}},
		{"'key' = 'value'", Config{Global: {"key": "value"}}},
		{"'key'=value; comment", Config{Global: {"key": "value"}}}, // Single qoute with comment.
		{"'key'=value ; comment", Config{Global: {"key": "value"}}},
		{"'key' = value; comment", Config{Global: {"key": "value"}}},
		{"'key' = value ; comment", Config{Global: {"key": "value"}}},
		{"key='value'; comment", Config{Global: {"key": "value"}}},
		{"key='value' ; comment", Config{Global: {"key": "value"}}},
		{"key = 'value'; comment", Config{Global: {"key": "value"}}},
		{"key = 'value' ; comment", Config{Global: {"key": "value"}}},
		{"'key'='value'; comment", Config{Global: {"key": "value"}}},
		{"'key'='value' ; comment", Config{Global: {"key": "value"}}},
		{"'key' = 'value'; comment", Config{Global: {"key": "value"}}},
		{"'key' = 'value' ; comment", Config{Global: {"key": "value"}}},
		{"'key'=value# comment", Config{Global: {"key": "value"}}},
		{"'key'=value # comment", Config{Global: {"key": "value"}}},
		{"'key' = value# comment", Config{Global: {"key": "value"}}},
		{"'key' = value # comment", Config{Global: {"key": "value"}}},
		{"key='value'# comment", Config{Global: {"key": "value"}}},
		{"key='value' # comment", Config{Global: {"key": "value"}}},
		{"key = 'value'# comment", Config{Global: {"key": "value"}}},
		{"key = 'value' # comment", Config{Global: {"key": "value"}}},
		{"'key'='value'# comment", Config{Global: {"key": "value"}}},
		{"'key'='value' # comment", Config{Global: {"key": "value"}}},
		{"'key' = 'value'# comment", Config{Global: {"key": "value"}}},
		{"'key' = 'value' # comment", Config{Global: {"key": "value"}}},
		{"'key'=; comment", Config{Global: {"key": ""}}}, // Single qoute only comment.
		{"'key'= ; comment", Config{Global: {"key": ""}}},
		{"'key' = ; comment", Config{Global: {"key": ""}}},
		{"key=''; comment", Config{Global: {"key": ""}}},
		{"key='' ; comment", Config{Global: {"key": ""}}},
		{"key = ''; comment", Config{Global: {"key": ""}}},
		{"key = '' ; comment", Config{Global: {"key": ""}}},
		{"'key'=''; comment", Config{Global: {"key": ""}}},
		{"'key'='' ; comment", Config{Global: {"key": ""}}},
		{"'key' = ''; comment", Config{Global: {"key": ""}}},
		{"'key' = '' ; comment", Config{Global: {"key": ""}}},
		{"'key'=# comment", Config{Global: {"key": ""}}},
		{"'key'= # comment", Config{Global: {"key": ""}}},
		{"'key' = # comment", Config{Global: {"key": ""}}},
		{"key=''# comment", Config{Global: {"key": ""}}},
		{"key='' # comment", Config{Global: {"key": ""}}},
		{"key = ''# comment", Config{Global: {"key": ""}}},
		{"key = '' # comment", Config{Global: {"key": ""}}},
		{"'key'=''# comment", Config{Global: {"key": ""}}},
		{"'key'='' # comment", Config{Global: {"key": ""}}},
		{"'key' = ''# comment", Config{Global: {"key": ""}}},
		{"'key' = '' # comment", Config{Global: {"key": ""}}},
		{"'key'=;", Config{Global: {"key": ""}}}, // Single quote empty comment.
		{"'key'= ;", Config{Global: {"key": ""}}},
		{"'key' = ;", Config{Global: {"key": ""}}},
		{"key='';", Config{Global: {"key": ""}}},
		{"key='' ;", Config{Global: {"key": ""}}},
		{"key = '';", Config{Global: {"key": ""}}},
		{"key = '' ;", Config{Global: {"key": ""}}},
		{"'key'='';", Config{Global: {"key": ""}}},
		{"'key'='' ;", Config{Global: {"key": ""}}},
		{"'key' = '';", Config{Global: {"key": ""}}},
		{"'key' = '' ;", Config{Global: {"key": ""}}},
		{`' key ' = ' value '`, Config{Global: {" key ": " value "}}}, // Keep whitespace.
		{`'	key	' = '	value	'`, Config{Global: {"	key	": "	value	"}}},
	}

	if err := testParser(tests); err != nil {
		t.Fatal(err.Error())
	}
}

func TestParseKeyValueMisc(t *testing.T) {
	t.Parallel()
	tests := []ParseTest{
		{`"=key"=value`, Config{Global: {"=key": "value"}}}, // Escaped qoutes.
		{`"k\"ey"=value`, Config{Global: {`k"ey`: "value"}}},
		{`key="val\"ue="`, Config{Global: {"key": `val"ue=`}}},
		{`"продавливания" = информации`, Config{Global: {"продавливания": `информации`}}}, // Issue #14.
		{"ke;y=value", Config{Global: {"ke;y": "value"}}},                                 // Misc.
		{"ke#y=value", Config{Global: {"ke#y": "value"}}},
		{`k\\ey=val\\ue`, Config{Global: {`k\ey`: `val\ue`}}},
		{`k\"ey=val\"ue`, Config{Global: {`k"ey`: `val"ue`}}},
		{`key=val\"ue\"`, Config{Global: {"key": `val"ue"`}}},
		{`key="val\"ue\""`, Config{Global: {"key": `val"ue"`}}},
		{`\\key=value`, Config{Global: {`\key`: "value"}}},
		{`"ke;y"=value`, Config{Global: {"ke;y": "value"}}},
		{`key="val;ue"`, Config{Global: {"key": "val;ue"}}},
		{`key="val#ue"`, Config{Global: {"key": "val#ue"}}},
		{`"ke;y"="val;ue"`, Config{Global: {"ke;y": "val;ue"}}},
		{`"ke#y"="val#ue"`, Config{Global: {"ke#y": "val#ue"}}},
		{`key==value`, Config{Global: {"key": "=value"}}},
		{`key=value=`, Config{Global: {"key": "value="}}},
	}

	if err := testParser(tests); err != nil {
		t.Fatal(err.Error())
	}
}

func TestParseSection(t *testing.T) {
	t.Parallel()
	tests := []ParseTest{
		{"[section]", Config{Global: {}, "section": {}}},
		{"[section];comment", Config{Global: {}, "section": {}}},
		{"[section] ; comment", Config{Global: {}, "section": {}}},
		{"[section]#comment", Config{Global: {}, "section": {}}},
		{"[section] # comment", Config{Global: {}, "section": {}}},
		{"[sec;tion]", Config{Global: {}, "sec;tion": {}}},
		{"[ s e c t i o n ]", Config{Global: {}, "s e c t i o n": {}}},
	}

	if err := testParser(tests); err != nil {
		t.Fatal(err.Error())
	}
}

func testParser(tests []ParseTest) error {
	for _, test := range tests {
		config, err := Parse(strings.NewReader(test.content))
		if err != nil {
			return fmt.Errorf("Unexpected error from Parse(%s): %s", test.content, err.Error())
		}

		if !reflect.DeepEqual(config, test.config) {
			return fmt.Errorf("Expected Parse(%s) to return %q, but got %q",
				test.content, test.config, config)
		}
	}

	return nil
}

func TestParseError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		content string
		errMsg  string
	}{
		{"key=value\nkey=value2", `ini: syntax error on line 2: ` +
			`key "key" already used in section "global"`},
		{"=value", `ini: syntax error on line 1: key can't be empty`},
		{`"key'=value`, `ini: syntax error on line 1: quote not closed`},
		{`"key=value`, `ini: syntax error on line 1: quote not closed`},
		{`'key"=value`, `ini: syntax error on line 1: quote not closed`},
		{`'key=value`, `ini: syntax error on line 1: quote not closed`},
		{`key="value'`, `ini: syntax error on line 1: quote not closed`},
		{`key="value`, `ini: syntax error on line 1: quote not closed`},
		{`key='value"`, `ini: syntax error on line 1: quote not closed`},
		{`key='value`, `ini: syntax error on line 1: quote not closed`},
		{"key", "ini: syntax error on line 1: no separator found"},
		{"key ; comment", "ini: syntax error on line 1: no separator found"},
		{"key # comment", "ini: syntax error on line 1: no separator found"},
		{"key value", "ini: syntax error on line 1: no separator found"},
		{`"key"`, `ini: syntax error on line 1: no separator found`},
		{`"key"value`, `ini: syntax error on line 1: unexpected "v", expected the separator "="`},
		{`"key"val=ue`, `ini: syntax error on line 1: unexpected "v", expected the separator "="`},
		{`"key" "value"`, `ini: syntax error on line 1: unexpected "\"", expected the separator "="`},
		{`"key" "2" = value`, `ini: syntax error on line 1: unexpected "\"", expected the separator "="`},
		{`"продавливания" информации`, `ini: syntax error on line 1: unexpected "и", expected the separator "="`}, // Issue #14.
		{"=value", "ini: syntax error on line 1: key can't be empty"},
		{"[", "ini: syntax error on line 1: unclosed section"},
		{"[section", `ini: syntax error on line 1: unclosed section`},
		{"[section] something", "ini: syntax error on line 1: unexpected \"s\" after section closed"},
		{"[section] информации", "ini: syntax error on line 1: unexpected \"и\" after section closed"},
		{"[]", "ini: syntax error on line 1: section can't be empty"},
		{"[ ]", "ini: syntax error on line 1: section can't be empty"},
		{"[section]\n[section]", "ini: syntax error on line 2: section \"section\" already exists"},
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
		} else if !IsSyntaxError(err) {
			t.Fatalf("Expected parseSection(%s) to return an syntax error, but it isn't",
				test.content)
		}
	}
}

func TestParseIOError(t *testing.T) {
	t.Parallel()
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
