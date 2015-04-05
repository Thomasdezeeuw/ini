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

func TestIndexUnescapedQouted(t *testing.T) {
	tests := []struct {
		input    string
		quote    string
		expected int
	}{
		{"key", "\"", -1},
		{"\"key", "\"", 0},
		{"'key", "'", 0},
		{"key\"", "\"", 3},
		{"key'", "'", 3},
		{"key\\\"", "\"", -1},
		{"key\\'", "'", -1},
		{"key\\\"\"", "\"", 5},
		{"key\\''", "'", 5},
	}

	for _, test := range tests {
		if got := indexUnescapedQouted(test.input, test.quote); got != test.expected {
			t.Fatalf("Expected indexUnqouted(%q) to return %d, but got %d",
				test.input, test.expected, got)
		}
	}
}

func TestGetKeyValue(t *testing.T) {
	tests := []struct {
		input string
		key   string
		value string
		err   error
	}{
		{"key=value", "key", "value", nil}, // Simple.
		{"k e y=v a l u e", "k e y", "v a l u e", nil},
		{"key = value", "key", "value", nil},
		{"key=", "key", "", nil},
		{"key=value; comment", "key", "value", nil}, // Simple with comment.
		{"key=value ; comment", "key", "value", nil},
		{"key = value; comment", "key", "value", nil},
		{"key = value ; comment", "key", "value", nil},
		{"key=; comment", "key", "", nil}, // Simple only comment.
		{"key= ; comment", "key", "", nil},
		{"key=;", "key", "", nil}, // Simple empty comment.
		{"key= ;", "key", "", nil},
		{"key = ;", "key", "", nil},
		{"key =  ;", "key", "", nil},

		{`"key"=value`, "key", "value", nil}, // Double qoute.
		{`"key" = value`, "key", "value", nil},
		{`key="value"`, "key", "value", nil},
		{`key = "value"`, "key", "value", nil},
		{`"key"="value"`, "key", "value", nil},
		{`"key" = "value"`, "key", "value", nil},
		{`"key"=value; comment`, "key", "value", nil}, // Double qoute with comment.
		{`"key"=value ; comment`, "key", "value", nil},
		{`"key" = value; comment`, "key", "value", nil},
		{`"key" = value ; comment`, "key", "value", nil},
		{`key="value"; comment`, "key", "value", nil},
		{`key="value" ; comment`, "key", "value", nil},
		{`key = "value"; comment`, "key", "value", nil},
		{`key = "value" ; comment`, "key", "value", nil},
		{`"key"="value"; comment`, "key", "value", nil},
		{`"key"="value" ; comment`, "key", "value", nil},
		{`"key" = "value"; comment`, "key", "value", nil},
		{`"key" = "value" ; comment`, "key", "value", nil},
		{`"key"=; comment`, "key", "", nil}, // Double qoute only comment.
		{`"key"= ; comment`, "key", "", nil},
		{`"key" = ; comment`, "key", "", nil},
		{`key=""; comment`, "key", "", nil},
		{`key="" ; comment`, "key", "", nil},
		{`key = ""; comment`, "key", "", nil},
		{`key = "" ; comment`, "key", "", nil},
		{`"key"=""; comment`, "key", "", nil},
		{`"key"="" ; comment`, "key", "", nil},
		{`"key" = ""; comment`, "key", "", nil},
		{`"key" = "" ; comment`, "key", "", nil},
		{`"key"=;`, "key", "", nil}, // Double quote empty comment.
		{`"key"= ;`, "key", "", nil},
		{`"key" = ;`, "key", "", nil},
		{`key="";`, "key", "", nil},
		{`key="" ;`, "key", "", nil},
		{`key = "";`, "key", "", nil},
		{`key = "" ;`, "key", "", nil},
		{`"key"="";`, "key", "", nil},
		{`"key"="" ;`, "key", "", nil},
		{`"key" = "";`, "key", "", nil},
		{`"key" = "" ;`, "key", "", nil},

		{"'key'=value", "key", "value", nil}, // Single qoute.
		{"'key' = value", "key", "value", nil},
		{"key='value'", "key", "value", nil},
		{"key = 'value'", "key", "value", nil},
		{"'key'='value'", "key", "value", nil},
		{"'key' = 'value'", "key", "value", nil},
		{"'key'=value; comment", "key", "value", nil}, // Single qoute with comment.
		{"'key'=value ; comment", "key", "value", nil},
		{"'key' = value; comment", "key", "value", nil},
		{"'key' = value ; comment", "key", "value", nil},
		{"key='value'; comment", "key", "value", nil},
		{"key='value' ; comment", "key", "value", nil},
		{"key = 'value'; comment", "key", "value", nil},
		{"key = 'value' ; comment", "key", "value", nil},
		{"'key'='value'; comment", "key", "value", nil},
		{"'key'='value' ; comment", "key", "value", nil},
		{"'key' = 'value'; comment", "key", "value", nil},
		{"'key' = 'value' ; comment", "key", "value", nil},
		{"'key'=; comment", "key", "", nil}, // Single qoute only comment.
		{"'key'= ; comment", "key", "", nil},
		{"'key' = ; comment", "key", "", nil},
		{"key=''; comment", "key", "", nil},
		{"key='' ; comment", "key", "", nil},
		{"key = ''; comment", "key", "", nil},
		{"key = '' ; comment", "key", "", nil},
		{"'key'=''; comment", "key", "", nil},
		{"'key'='' ; comment", "key", "", nil},
		{"'key' = ''; comment", "key", "", nil},
		{"'key' = '' ; comment", "key", "", nil},
		{"'key'=;", "key", "", nil}, // Single quote empty comment.
		{"'key'= ;", "key", "", nil},
		{"'key' = ;", "key", "", nil},
		{"key='';", "key", "", nil},
		{"key='' ;", "key", "", nil},
		{"key = '';", "key", "", nil},
		{"key = '' ;", "key", "", nil},
		{"'key'='';", "key", "", nil},
		{"'key'='' ;", "key", "", nil},
		{"'key' = '';", "key", "", nil},
		{"'key' = '' ;", "key", "", nil},

		// todo: below

		{`"=key"=value`, "=key", "value", nil}, // Escaped qoutes.
		{`"k\"ey"=value`, `k"ey`, "value", nil},
		{`key="val\"ue="`, "key", `val"ue=`, nil},

		{"ke;y=value", "ke;y", "value", nil}, // Misc.
		{`"ke;y"=value`, "ke;y", "value", nil},
		{`key="val;ue"`, "key", "val;ue", nil},
		{`"ke;y"="val;ue"`, "ke;y", "val;ue", nil},
		{`key==value`, "key", `=value`, nil},
		{`key=value=`, "key", `value=`, nil},

		{`"key'=value`, "", "", errSynthax}, // Incorret.
		{`key="value'`, "", "", errSynthax},
		{`'key"=value`, "", "", errSynthax},
		{`key='value"`, "", "", errSynthax},
		{"key", "", "", errSynthax},
		{"key value", "", "", errSynthax},
		{`"key"`, "", "", errSynthax},
		{`"key"value`, "", "", errSynthax},
		{`"key"val=ue`, "", "", errSynthax},
		{"=value", "", "", errSynthax},
	}

	for _, test := range tests {
		key, value, err := getKeyValue(test.input)
		if err != test.err {
			t.Fatalf("Expected getKeyValue(%q) to return error \"%v\", got \"%v\"",
				test.input, test.err, err)
		} else if key != test.key {
			t.Fatalf("Expected getKeyValue(%q) to return key %q, but got %q",
				test.input, test.key, key)
		} else if value != test.value {
			t.Fatalf("Expected getKeyValue(%q) to return value %q, but got %q",
				test.input, test.value, value)
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
		{"testdata/malformed.cfg", nil, errSynthax, false},
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
			if err != test.err {
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

	for section, section1 := range c {
		for key, value1 := range section1 {
			value2 := c2[section][key]
			if value1 != value2 {
				t.Fatalf("Expected section %q to key %q to be the same, but got %q and %q",
					section, key, value1, value2)
			}
		}
	}

	// Cleanup
	if err := os.Remove(tmpPath); err != nil {
		t.Fatal(err)
	}
}
