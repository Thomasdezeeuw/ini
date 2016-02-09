// Copyright (C) 2015-2016 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

// Package ini parses an ini formatted file. The parsed configuration can be
// decoded into a variable or used directly.
package ini

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
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
type Config map[string]Section

// Section holds the key value pairs inside a configuration.
//
//	section := config["section"]
//	value, found := section["key"]
type Section map[string]string

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
		keys := getSectionKeysAlpha(section)
		for _, key := range keys {
			value := strconv.Quote(section[key])
			key = strconv.Quote(key)
			result.WriteString(key + "=" + value + "\n")
		}
		result.WriteString("\n")
	}

	return &result
}

// GetConfigSectionsAlpha sorts the sections alphabetically with the global
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

// GetSectionKeysAlpha sorts the keys in a section alphabetically.
func getSectionKeysAlpha(m Section) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

type fieldCombo struct {
	value reflect.Value
	field reflect.StructField
}

// Decode decodes a configuration into a struct. Any properties to be set need
// to be public. Keys are renamed, whitespace is removed and keys start with a
// capaital, like so:
//
//	"my key" -> "MyKey"
//
// Tags are supported to define the key of a configuration. See below for a
// small example, for more examples see the tags example in the _examples
// directory.
//
//	struct {
//		AppName `ini:"name"`
//	}
//
// Slices are supported by using a comma separated list, like so:
//
//	"string1, string2" -> []string{"string1", "string2"}
//	"1, 2, 3" -> []int{1, 2, 3}
//
// Booleans are supported as well:
//
//	"1, t, T, TRUE, true, True" -> true
//	"0, f, F, FALSE, false, False" -> false
//
// Time is supported with the following formats:
//
//	"2006-01-02", "2006-01-02 15:04", "2006-01-02 15:04:05", time.RFC3339,
//	time.RFC1123 and time.RFC822
//
// Duration is also supported, see `time.ParseDuration` for the documentation.
//
// Note: underneath Decode uses the reflect package which isn't great for
// performance, so use it with care.
func (c *Config) Decode(dst interface{}) error {
	valuePtr := reflect.ValueOf(dst)
	value := reflect.Indirect(valuePtr)

	// If it's not a pointer we can't change the original value and if it's not a
	// struct we can't set/change any keys on it, in either case we can't do
	// anything with the value.
	if valuePtr.Kind() != reflect.Ptr || value.Kind() != reflect.Struct {
		return errors.New("ini: Config.Decode requires a pointer to a struct")
	}

	dstType := value.Type()
	for i := value.NumField() - 1; i >= 0; i-- {
		field := value.Field(i)
		if !field.IsValid() || !field.CanSet() {
			continue
		}

		strField := dstType.Field(i)
		var fields []fieldCombo
		var sectionNames = []string{""}

		structFieldType := field.Type()
		if kind := field.Kind(); kind == reflect.Struct &&
			structFieldType != typeDuration && structFieldType != typeTime {
			sectionName := strField.Tag.Get("ini")
			if sectionName != "" {
				sectionNames[0] = sectionName
			} else {
				sectionNames = possibleNames(strField.Name)
			}

			for i := field.NumField() - 1; i >= 0; i-- {
				structField := field.Field(i)
				if !structField.IsValid() || !structField.CanSet() {
					continue
				}
				structFieldType := structFieldType.Field(i)
				fields = append(fields, fieldCombo{structField, structFieldType})
			}
		} else {
			fields = []fieldCombo{{field, strField}}
		}

		for _, combo := range fields {
			value := combo.value
			field := combo.field

			var keys []string
			if key := field.Tag.Get("ini"); key != "" {
				keys = []string{key}
			} else {
				keys = possibleNames(field.Name)
			}

			if err := c.trySetReflect(sectionNames, keys, value); err != nil {
				return err
			}
		}
	}

	return nil
}

// TrySetReflect tries the givens section and keys combination to get the value
// from the config and then sets the field if a value if found.
func (c *Config) trySetReflect(sectionNames []string, keys []string, keyValue reflect.Value) error {
	for _, sectionName := range sectionNames {
		section := (*c)[sectionName]
		if len(section) == 0 {
			continue
		}

		for _, key := range keys {
			value, ok := section[key]
			if !ok {
				continue
			}

			if err := setReflectValue(&keyValue, value); err != nil {
				if sectionName == Global {
					sectionName = globalName
				}
				return fmt.Errorf("ini: error decoding %q in section %q: %s",
					key, sectionName, err.Error())
			}

			return nil
		}
	}

	return nil
}

var separators = []string{"", "_", "-", " "}

// PossibleNames generates possible names for a given name. It splits up the
// given name on upper case and underscores (see getNameParts) and then joins
// them using nothing (""), underscores, dashes and spaces (see separators
// variable).
func possibleNames(name string) []string {
	nameParts := getNameParts(name)
	if len(nameParts) == 1 {
		return []string{name, strings.ToLower(name)}
	}

	var names []string
	for _, seperator := range separators {
		var newName string
		for _, namePart := range nameParts {
			newName += namePart
			newName += seperator
		}
		newName = newName[:len(newName)-len(seperator)]

		names = append(names, strings.ToLower(newName))
		names = append(names, newName)
	}

	return names
}

// GetNameParts splits a name on upper case and underscores, it returns the
// split parts.
// todo: maybe split on first number?
func getNameParts(name string) []string {
	if name == "" {
		return []string{""}
	}

	var nameParts = []string{name[:1]}
	var i int
	for _, r := range name[1:] {
		if r == '_' || unicode.IsUpper(r) {
			nameParts = append(nameParts, string(r))
			i++
		} else {
			nameParts[i] += string(r)
		}
	}

	return nameParts
}
