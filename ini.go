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
	"reflect"
	"sort"
	"strings"
)

// Global is the section name for key-values not under a section. It is used
// for retrieving key-value pairs in the "global" section, or not under any
// other section.
//
//	value := config[ini.Global]["key"]
//	value2, found := config[ini.Global]["key2"]
const Global = ""

// Name for the global section in errors.
const globalName = "global"

// Config holds all key-value pairs under sections. To retrieve a key-value
// pair you can use:
//
//	value := config["section"]["key"]
//	value2, found := config["section"]["key2"]
//
// The same as you would for a map. For retrieving key-value pairs not in any
// section, or the "global" section, you can use the `ini.Global` constant,
// like so:
//
//	value := config[ini.Global]["key"]
//	value2, found := config[ini.Global]["key2"]
type Config map[string]map[string]string

// String returns an ini formatted configuration, ready to be written to a file.
func (c *Config) String() string {
	return c.buffer().String()
}

// Bytes returns an ini formatted configuration, ready to be written to a file.
func (c *Config) Bytes() []byte {
	return c.buffer().Bytes()
}

// WriteTo writes the configuration to the writer in the ini format.
func (c *Config) WriteTo(w io.Writer) (int64, error) {
	return c.buffer().WriteTo(w)
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
			value := possibleQoute(section[key])
			key = possibleQoute(key)

			result.WriteString(key + "=" + value + "\n")
		}
		result.WriteString("\n")
	}

	return &result
}

// Scan scans a configuration into a struct. Any properties to be set need to
// be public. Keys are renamed, whitespace is removed and keys start with a
// capaital, like so:
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
// Duration is also supported, it uses `time.ParseDuration` so for supported
// format see the seperate documentation.
//
// *Note underneath Scan uses the reflect package which isn't great for
// performance, so use it with care.
func (c *Config) Scan(dst interface{}) error {
	valuePtr := reflect.ValueOf(dst)
	value := reflect.Indirect(valuePtr)

	// If it's not a pointer we can't change the original value and if it's not a
	// struct we can't set/change any keys on it, in either case we can't do
	// anything with the value.
	if valuePtr.Kind() != reflect.Ptr || value.Kind() != reflect.Struct {
		return errors.New("ini: Config.Scan requires a pointer to a struct")
	}

	for sectionName, section := range *c {
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
					sectionName = globalName
				}
				return fmt.Errorf("ini: error scanning %q in section %q: %s",
					key, sectionName, err.Error())
			}
		}
	}

	return nil
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

// PossibleQoute adds qouting to key or values when needed. For example:
//
//	`my key`   -> `my key`       // no qoutes.
//	`my "key"` -> `"my \"key\""` // qoutes and escaped qoutes.
//	`my 'key'` -> `"my 'key'"`   // just qoutes.
func possibleQoute(value string) string {
	if strings.Index(value, `"`) != -1 {
		return `"` + strings.Replace(value, `"`, `\"`, -1) + `"`
	} else if strings.Index(value, `'`) != -1 {
		return `"` + value + `"`
	}

	return value
}
