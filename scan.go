// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func setReflectValue(keyValue *reflect.Value, value string) error {
	// todo: improve the detection of time.Duration and time.Time.
	if tStr := fmt.Sprintf("%v", keyValue.Type()); tStr == "time.Duration" {
		return setDuration(keyValue, value)
	} else if tStr == "time.Time" {
		return setTime(keyValue, value)
	}

	switch keyValue.Kind() {
	case reflect.String:
		keyValue.SetString(value)
	case reflect.Bool:
		return setBool(keyValue, value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setInt(keyValue, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return setUint(keyValue, value)
	case reflect.Float32, reflect.Float64:
		return setFloat(keyValue, value)
	case reflect.Slice:
		return setSlice(keyValue, value)
	}

	return nil
}

// SetSlice sets a slice of reflected value.
func setSlice(f *reflect.Value, value string) error {
	if !f.IsValid() || !f.CanSet() {
		return nil
	}

	values := strings.Split(value, ",")
	for i, value := range values {
		values[i] = strings.TrimSpace(value)
	}

	// todo: clean up switch
	switch f.Type().Elem().Kind() {
	case reflect.String:
		f.Set(reflect.ValueOf(values))
	case reflect.Bool:
		var bs []bool
		for _, value := range values {
			bs = append(bs, parseBool(value))
		}
		f.Set(reflect.ValueOf(bs))
	case reflect.Int:
		var ints []int
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			ints = append(ints, i)
		}
		f.Set(reflect.ValueOf(ints))
	case reflect.Int8:
		var ints []int8
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			} else if i > math.MaxInt8 || i < math.MinInt8 {
				return fmt.Errorf("ini: %d overflows %q", i, f.Type().Elem().Kind())
			}
			ints = append(ints, int8(i))
		}
		f.Set(reflect.ValueOf(ints))
	case reflect.Int16:
		var ints []int16
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			} else if i > math.MaxInt16 || i < math.MinInt16 {
				return fmt.Errorf("ini: %d overflows %q", i, f.Type().Elem().Kind())
			}
			ints = append(ints, int16(i))
		}
		f.Set(reflect.ValueOf(ints))
	case reflect.Int32:
		var ints []int32
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			} else if i > math.MaxInt32 || i < math.MinInt32 {
				return fmt.Errorf("ini: %d overflows %q", i, f.Type().Elem().Kind())
			}
			ints = append(ints, int32(i))
		}
		f.Set(reflect.ValueOf(ints))
	case reflect.Int64:
		var ints []int64
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			ints = append(ints, int64(i))
		}
		f.Set(reflect.ValueOf(ints))
	case reflect.Uint:
		var is []uint
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			} else if i < 0 {
				return fmt.Errorf("ini: %d overflows %q", i, f.Type().Elem().Kind())
			}
			is = append(is, uint(i))
		}
		f.Set(reflect.ValueOf(is))
	case reflect.Uint8:
		var is []uint8
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			} else if i > math.MaxUint8 || i < 0 {
				return fmt.Errorf("ini: %d overflows %q", i, f.Type().Elem().Kind())
			}
			is = append(is, uint8(i))
		}
		f.Set(reflect.ValueOf(is))
	case reflect.Uint16:
		var is []uint16
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			} else if i > math.MaxUint16 || i < 0 {
				return fmt.Errorf("ini: %d overflows %q", i, f.Type().Elem().Kind())
			}
			is = append(is, uint16(i))
		}
		f.Set(reflect.ValueOf(is))
	case reflect.Uint32:
		var is []uint32
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			} else if i > math.MaxUint32 || i < 0 {
				return fmt.Errorf("ini: %d overflows %q", i, f.Type().Elem().Kind())
			}
			is = append(is, uint32(i))
		}
		f.Set(reflect.ValueOf(is))
	case reflect.Uint64:
		var is []uint64
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			is = append(is, uint64(i))
		}
		f.Set(reflect.ValueOf(is))
	case reflect.Uintptr:
		var is []uintptr
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			} else if i > math.MaxUint8 || i < 0 {
				return fmt.Errorf("ini: %d overflows %q", i, f.Type().Elem().Kind())
			}
			is = append(is, uintptr(i))
		}
		f.Set(reflect.ValueOf(is))
	case reflect.Float32:
		var fs []float32
		for _, value := range values {
			fv, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			} else if fv > math.MaxFloat32 {
				return fmt.Errorf("ini: %f overflows %q", fv, f.Type().Elem().Kind())
			}
			fs = append(fs, float32(fv))
		}
		f.Set(reflect.ValueOf(fs))
	case reflect.Float64:
		var fs []float64
		for _, value := range values {
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fs = append(fs, f)
		}
		f.Set(reflect.ValueOf(fs))
	}

	return nil
}

func setDuration(keyValue *reflect.Value, value string) error {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return createConvertError(value, "time.Duration")
	}

	durationValue := reflect.ValueOf(duration)
	keyValue.Set(durationValue)
	return nil
}

func setTime(keyValue *reflect.Value, value string) error {
	for _, format := range timeFormats {
		t, err := time.Parse(format, value)
		if err == nil {
			timeValue := reflect.ValueOf(t)
			keyValue.Set(timeValue)
			return nil
		}
	}

	return createConvertError(value, "time.Time")
}

// Returns true on "1" and "true", anything returns false.
func setBool(keyValue *reflect.Value, value string) error {
	b := parseBool(value)
	keyValue.SetBool(b)
	return nil
}

func parseBool(value string) bool {
	var b bool
	value = strings.TrimSpace(value)
	if value == "1" || strings.ToLower(value) == "true" {
		b = true
	}
	return b
}

func setInt(keyValue *reflect.Value, value string) error {
	i, err := strconv.Atoi(value)
	if err != nil {
		return createConvertError(value, keyValue.Kind().String())
	}
	i64 := int64(i)

	if keyValue.OverflowInt(i64) {
		createOverflowError(value, keyValue.Kind().String())
	}

	keyValue.SetInt(i64)
	return nil
}

func setUint(keyValue *reflect.Value, value string) error {
	i, err := strconv.Atoi(value)
	if err != nil {
		return createConvertError(value, keyValue.Kind().String())
	}
	ui64 := uint64(i)

	if keyValue.OverflowUint(ui64) {
		createOverflowError(value, keyValue.Kind().String())
	}

	keyValue.SetUint(ui64)
	return nil
}

func setFloat(keyValue *reflect.Value, value string) error {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return createConvertError(value, keyValue.Kind().String())
	}

	if keyValue.OverflowFloat(f) {
		createOverflowError(value, keyValue.Kind().String())
	}

	keyValue.SetFloat(f)
	return nil
}

// Rename the key to a public name of a struct, e.g.
//
//	"my key" -> "MyKey"
func renameToPublicName(name string) string {
	name = strings.Title(name)
	name = strings.Replace(name, " ", "", -1)
	return name
}

func getSectionValue(keyValue reflect.Value, sectionName string) reflect.Value {
	if sectionName == Global {
		return keyValue
	}

	sectionName = renameToPublicName(sectionName)
	return keyValue.FieldByName(sectionName)
}

func createOverflowError(value, t string) error {
	return fmt.Errorf("can't convert %q to type %s, it overflows type %s",
		value, t, t)
}

func createConvertError(value, t string) error {
	return fmt.Errorf("can't convert %q to type %s", value, t)
}
