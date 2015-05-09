// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

// Package ini parses an ini formatted file. The parsed configuration can be
// scanned into a variable or used directly.
package ini

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
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

var timeFormats = []string{"2006-01-02", "2006-01-02 15:04",
	"2006-01-02 15:04:05"}

// Config holds all key-value pairs under sections. To retrieve keys use:
//
//	value, found := config["section"]["key"]
//
// You can use the `ini.Global` constant to retrieve key-value pairs not under
// any section, or the "global" section:
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

// Scan scans a configuration into a struct or map. Any properties to be set
// need to be public. Keys are renamed, whitespace is removed and keys start
// with a capaital, like so:
//
//	"my key" -> "MyKey"
//
// Slices are supported by using a comma separated list, like so:
//
//	"string1, string2" -> []string{"string1", "string2"}
//	"1, 2, 3" -> []int{1, 2, 3}
//
// Booleans are supported aswell:
//
//	"1, true, TRUE" -> true
//	"0, false, FALSE, anything else" -> false
//
// Time is supported with the following formats:
//
//	"2006-01-02", "2006-01-02 15:04", "2006-01-02 15:04:05"
//
// *Note underneath Scan uses the reflect package which isn't great for
// performance, so use it with care.
func (c Config) Scan(dst interface{}) error {
	valuePtr := reflect.ValueOf(dst)
	value := reflect.Indirect(valuePtr)

	// If it's not a pointer we can't change the original value and if it's not a
	// struct we can't set/change any keys on it, in either case we can't do
	// anything with the value.
	if valuePtr.Kind() != reflect.Ptr || value.Kind() != reflect.Struct {
		return errors.New("ini: ini.Config.Scan requires a pointer to struct")
	}

	for sectionName, section := range c {
		sectionValue := getSectionValue(value, sectionName)
		if !sectionValue.IsValid() || sectionValue.Kind() != reflect.Struct {
			continue
		}

		for key, value := range section {
			key = renameToPublicName(key)

			keyValue := sectionValue.FieldByName(key)
			if !keyValue.IsValid() || !keyValue.CanSet() {
				continue
			}

			if err := setReflectValue(&keyValue, value); err != nil {
				if sectionName == Global {
					sectionName = "Global"
				}
				return fmt.Errorf("ini: error scanning %q in section %q: %s",
					key, sectionName, err.Error())
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

// GetConfigSectionsAlpha sort the sections alphabetically with the global
// section first.
func getConfigSectionsAlpha(c Config) []string {
	sections := make([]string, 0, len(c))
	for section := range c {
		if section != Global {
			sections = append(sections, section)
		}
	}
	sort.Strings(sections)
	return append([]string{Global}, sections...)
}
