// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import (
	"fmt"
	"math"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	typeString   = reflect.TypeOf("")
	typeBool     = reflect.TypeOf(true)
	typeInt      = reflect.TypeOf(int(1))
	typeInt8     = reflect.TypeOf(int8(1))
	typeInt16    = reflect.TypeOf(int16(1))
	typeInt32    = reflect.TypeOf(int32(1))
	typeInt64    = reflect.TypeOf(int64(1))
	typeUint     = reflect.TypeOf(uint(1))
	typeUint8    = reflect.TypeOf(uint8(1))
	typeUint16   = reflect.TypeOf(uint16(1))
	typeUint32   = reflect.TypeOf(uint32(1))
	typeUint64   = reflect.TypeOf(uint64(1))
	typeFloat32  = reflect.TypeOf(float32(1.0))
	typeFloat64  = reflect.TypeOf(float64(1.0))
	typeDuration = reflect.TypeOf(time.Nanosecond)
	typeTime     = reflect.TypeOf(time.Time{})
)

// Scan scans a configuration into a struct or map, see `Config.Scan`.
func Scan(r io.Reader, dst interface{}) error {
	c, err := Parse(r)
	if err != nil {
		return err
	}
	return c.Scan(dst)
}

func setReflectValue(keyValue *reflect.Value, value string) error {
	if keyValue.Kind() == reflect.Slice {
		return setSlice(keyValue, value)
	}

	// todo: (maybe) add map and struct (convert to a section maybe?).
	switch keyValue.Type() {
	case typeString:
		keyValue.SetString(value)
	case typeBool:
		return setBool(keyValue, value)
	case typeInt, typeInt8, typeInt16, typeInt32, typeInt64:
		return setInt(keyValue, value)
	case typeUint, typeUint8, typeUint16, typeUint32, typeUint64:
		return setUint(keyValue, value)
	case typeFloat32, typeFloat64:
		return setFloat(keyValue, value)
	case typeDuration:
		return setDuration(keyValue, value)
	case typeTime:
		return setTime(keyValue, value)
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

// Returns true on "1" and "true", anything returns false.
func setBool(keyValue *reflect.Value, value string) error {
	var b bool
	value = strings.TrimSpace(value)
	if value == "1" || strings.ToLower(value) == "true" {
		b = true
	}
	keyValue.SetBool(b)
	return nil
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

func getSectionValue(keyValue reflect.Value, sectionName string) reflect.Value {
	if sectionName == Global {
		return keyValue
	}

	sectionName = renameToPublicName(sectionName)
	return keyValue.FieldByName(sectionName)
}

// Rename the key to a public name of a struct, e.g.
//
//	"my key" -> "MyKey"
func renameToPublicName(name string) string {
	name = strings.Title(name)
	name = strings.Replace(name, " ", "", -1)
	return name
}

func getValues(value string) []string {
	values := strings.Split(value, ",")
	for i, value := range values {
		values[i] = strings.TrimSpace(value)
	}
	return values
}
