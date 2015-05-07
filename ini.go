// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

// Package ini parses an ini formatted file. It can load it into a struct or
// used directly.
package ini

import (
	"bytes"
	"errors"
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
	if vPtr.Kind() != reflect.Ptr || v.Kind() != reflect.Struct {
		return errors.New("ini: ini.Config.Scan requires a pointer to struct")
	}

	for sectionName, section := range c {
		var sf reflect.Value
		if sectionName == Global {
			sf = v
		} else {
			// Rename the section like the following: "my section" -> "MySection".
			sectionName = strings.Replace(strings.Title(sectionName), " ", "", -1)
			sf = v.FieldByName(sectionName)
		}

		// Make sure it's valid and a map or a struct, otherwise we can't do
		// anything with it.
		if !sf.IsValid() || sf.Kind() != reflect.Struct {
			continue
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

// Scan scans a configuration into a struct or map, see #Config.Scan.
func Scan(path string, dst interface{}) error {
	c, err := Load(path)
	if err != nil {
		return err
	}
	return c.Scan(dst)
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
