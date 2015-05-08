// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

// Package ini parses an ini formatted file. The parsed configuration can be
// scanned into a variable or used directly.
package ini

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

// Global is the section name for key-values not under a section. It is used
// for retrieving key-value pairs in the "global" section, or not under any
// other section.
//
//	value, found := config[ini.Global]["key"]
const Global = global

// We hide the public Globals value to hope people won't use the string instead
// of the constant.
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

// String returns an ini formatted configuration, ready to be written to a file.
func (c *Config) String() string {
	return c.buffer().String()
}

// Bytes returns an ini formatted configuration, ready to be written to a file.
func (c *Config) Bytes() []byte {
	return c.buffer().Bytes()
}

// Buffer creates a `bytes.Buffer` with an ini formatted configuration.
func (c *Config) buffer() *bytes.Buffer {
	var result bytes.Buffer

	sections := getConfigSectionsAlpha(*c)
	for _, sectionName := range sections {
		if sectionName != Global {
			result.WriteString("[" + sectionName + "]\n")
		}

		section := (*c)[sectionName]
		keys := getMapsKeysAlpha(section)
		for _, key := range keys {
			value := section[key]
			result.WriteString(key + "=" + value + "\n")
		}
		result.WriteString("\n")
	}

	return &result
}

// WriteTo writes the configuration to the writer in the ini format.
func (c *Config) WriteTo(w io.Writer) (int64, error) {
	return c.buffer().WriteTo(w)
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
		return nil, errors.New("ini: only loading files is supported")
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

func getMapsKeysAlpha(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func getConfigSectionsAlpha(c Config) []string {
	// Make sure the global section is the first section.
	sections := make([]string, 0, len(c)-1)
	for section := range c {
		if section != Global {
			sections = append(sections, section)
		}
	}
	sort.Strings(sections)
	return append([]string{Global}, sections...)
}
