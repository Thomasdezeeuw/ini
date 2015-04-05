// Package ini parses an ini formatted file. It can load it into a struct or
// used directly.
package ini

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

var (
	errFile = errors.New("ini: only loading files is supported")

	// todo: create a error type for synthax error, with line a column. Easier
	// for debugging.
	errSynthax = errors.New("ini: incorrect formatted synthax")
)

// Global is the section name for key-values not under a section. It is used
// for retrieving key-value pairs in the "global" section, or not under any
// other section.
//
//	value, found := config[ini.Global]["key"]
const Global = global
const global = "SUPERGLOBAL"

// Different parts of the ini format.
const (
	commentStart = ";"
	seperator    = "="
	sectionStart = "["
	sectionEnd   = "]"
	doubleQuote  = `"`
	singleQuote  = "'"
	escapeQuote  = `\`
)

// Config holds all key-value pairs under sections. To retrieve keys use:
//
//	value, found := config["section"]["key"]
//
// And use the `ini.Global` constant to retrieve key-value pairs not under any
// section, or the "global" section:
//
//	value, found := config[ini.Global]["key"]
type Config map[string]map[string]string

// String returns a ini-formatted configuration, ready to be written to a file.
func (c Config) String() string {
	var result bytes.Buffer

	// Order the sections alphabetically, ignore the "global" section.
	sections := make([]string, 0, len(c)-1)
	for section := range c {
		if section != Global {
			sections = append(sections, section)
		}
	}
	sort.Strings(sections)

	// Write each section starting with the "global" one.
	for _, section := range append([]string{Global}, sections...) {
		// Write the section header, except for the "global" section.
		if section != Global {
			result.WriteString("[" + section + "]\n")
		}

		// Write each key-value pair.
		confSection := c[section]
		for _, key := range getKeysAlpha(confSection) {
			result.WriteString(key + "=" + confSection[key] + "\n")
		}
		result.WriteString("\n")
	}

	return result.String()
}

// Scan scan a configuration into a struct or map. Any properties to be set
// need to be exported. Key are renamed, whitespace is removed and keys start
// with a capaital letter, like so:
//
//	"my key" -> "MyKey"
//
// Slices are exported by using a comma separated list, like so:
//
//	"string1, string2" -> []string{"string1", "string2"}
//	"1, 2, 3" -> []int{1, 2, 3}
//
// Booleans are supported aswell:
//
//	"1, true, TRUE" // true
//	"0, false, FALSE, anything else" // false
//
// *Note underneath Scan uses reflect which isn't great for performance, use it
// with care.
func (c Config) Scan(dst interface{}) error {
	vPtr := reflect.ValueOf(dst)
	v := reflect.Indirect(vPtr)

	// If it's not a pointer we can't change the value, so it's pointless to do
	// anything. If it's not a struct or we can't set any keys on it.
	if k := v.Kind(); vPtr.Kind() != reflect.Ptr || (k != reflect.Struct && k != reflect.Map) {
		return errors.New("ini: ini.Config.Scan requires a pointer to struct")
	}

	for sectionName, section := range c {
		var sf reflect.Value
		if sectionName == Global {
			sf = v
		} else {
			sf = v.FieldByName(sectionName)
		}

		// Make sure it's valid and a map or a struct, otherwise we can't do
		// anything with it.
		if k := sf.Kind(); !sf.IsValid() || (k != reflect.Struct && k != reflect.Map) {
			return fmt.Errorf("ini: no struct %q on desination", sectionName)
		}

		// Set each key value in the section.
		for key, value := range section {
			// Rename the key like the following: "my key" -> "MyKey".
			key = strings.Replace(strings.Title(key), " ", "", -1)

			// Try to get  the field and make sure we can change it.
			if f := sf.FieldByName(key); !f.IsValid() || !f.CanSet() {
				continue
			} else if err := setReflectValue(&f, value); err != nil {
				return err
			}
		}
	}

	return nil
}

// Load a single configuration file.
func Load(path string) (Config, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Make sure it's a file.
	if s, err := f.Stat(); err != nil {
		return nil, err
	} else if s.IsDir() {
		return nil, errFile
	}

	return Parse(f)
}

// Parse parses an ini format file.
func Parse(r io.Reader) (Config, error) {
	c := Config{}

	// Start in the global section.
	section := Global
	c[section] = map[string]string{}

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		// Ignore empty lines and comments.
		if len(line) == 0 || line[:1] == commentStart {
			continue
		}

		// Handle sections.
		if line[:1] == sectionStart && line[len(line)-1:] == sectionEnd {
			section = line[1 : len(line)-1]
			c[section] = map[string]string{}
			continue
		}

		// Handle key-value pairs.
		key, value, err := getKeyValue(line)
		if err != nil {
			return nil, err
		}
		c[section][key] = value
	}

	return c, nil
}

// Scan scans a configuration into a struct or map, see #Config.Scan.
func Scan(path string, dst interface{}) error {
	c, err := Load(path)
	if err != nil {
		return err
	}
	return c.Scan(dst)
}

// GetKeyValue gets a key value pair from a single line.
func getKeyValue(line string) (key string, value string, err error) {
	line = strings.TrimSpace(line)
	// No seperator means no key-value pair, so a synthax error.
	if count := strings.Count(line, seperator); count == 0 {
		return "", "", errSynthax
	}

	// Determine if the key is qouted or not.
	if quote := line[:1]; quote == doubleQuote || quote == singleQuote {
		i := indexUnescapedQouted(line[1:], quote)
		// No more unescaped qoutes found.
		if i == -1 {
			return "", "", errSynthax
		}
		i++ // we cut of the first quote.

		// Skip the first quoute and include up to the next unescaped qoute.
		key = strings.Replace(line[1:i], escapeQuote+quote, quote, -1)

		// For the value skip the key and it's last unescaped qoute, trim it and
		// the seperator should follow next, if not it's incorrect synthax. Finally
		// drop the seperator.
		value = strings.TrimSpace(line[i+1:])
		if string(value[0]) != seperator {
			return "", "", errSynthax
		}
		value = value[1:]
	} else {
		// Simple split on the seperator.
		v := strings.SplitN(line, seperator, 2)
		key, value = v[0], v[1]
	}

	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	// Key can't be empty, otherwise we can't put it in the map.
	if len(key) == 0 {
		return "", "", errSynthax
	}

	if len(value) != 0 {
		// If the value is just a comment start, remove it. Otherwise check if the
		// value is qouted or has a comment added to it.
		if len(value) == 1 && value == commentStart {
			value = ""
		} else if quote := value[:1]; quote == doubleQuote || quote == singleQuote {
			i := indexUnescapedQouted(value[1:], quote)
			// No more unescaped qoutes found
			if i == -1 {
				return "", "", errSynthax
			}
			i++ // we cut of the first quote

			// Skip the first quoute and include up to the next unescaped qoute.
			value = strings.Replace(value[1:i], escapeQuote+quote, quote, -1)
			value = strings.TrimSpace(value)
		} else if i := strings.Index(value, commentStart); i != -1 {
			// Drop any comment attached to the value
			value = strings.TrimSpace(value[:i])
		}
	}

	return key, value, nil
}

// indexUnescapedQouted gets the index of the next unescaped qoute.
func indexUnescapedQouted(line, quote string) int {
	i := strings.Index(line, quote)
	for lineLeft := line; i != -1; i = strings.Index(lineLeft, quote) {
		// If the quote is not escaped it is the closing quote and the key
		// ends here.
		if i == 0 || string(lineLeft[i-1]) != escapeQuote {
			// The end is the index of the found quoute plus the difference in
			// the length of the line and the length of what is left of the line
			// plus 1 (0 index, want to include the quote character).
			return i + (len(line) - len(lineLeft))
		}

		// Escaped quote, on to the next one.
		lineLeft = lineLeft[i+1:]
	}

	// No unqouted qoute found
	return -1
}

// GetKeysAlpha returns keys of the map sorted alphabetically.
func getKeysAlpha(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
