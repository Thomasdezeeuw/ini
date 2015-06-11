// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import (
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var timeFormats = []string{"2006-01-02", "2006-01-02 15:04",
	"2006-01-02 15:04:05"}

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

	switch keyValue.Type() {
	case typeString:
		keyValue.SetString(value)
	case typeBool:
		setBool(keyValue, value)
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
func setSlice(keyValue *reflect.Value, value string) error {
	values := getValues(value)

	// Type of the slice elements.
	switch keyValue.Type().Elem() {
	case typeString:
		keyValue.Set(reflect.ValueOf(values))
	case typeBool:
		return setBools(keyValue, values)
	case typeInt:
		return setInts(keyValue, values)
	case typeInt8:
		return setInt8s(keyValue, values)
	case typeInt16:
		return setInt16s(keyValue, values)
	case typeInt32:
		return setInt32s(keyValue, values)
	case typeInt64:
		return setInt64s(keyValue, values)
	case typeUint:
		return setUints(keyValue, values)
	case typeUint8:
		return setUint8s(keyValue, values)
	case typeUint16:
		return setUint16s(keyValue, values)
	case typeUint32:
		return setUint32s(keyValue, values)
	case typeUint64:
		return setUint64s(keyValue, values)
	case typeFloat32:
		return setFloat32s(keyValue, values)
	case typeFloat64:
		return setFloat64s(keyValue, values)
	case typeDuration:
		return setDurations(keyValue, values)
	case typeTime:
		return setTimes(keyValue, values)
	}

	return nil
}

func setBools(keyValue *reflect.Value, values []string) error {
	var bs = make([]bool, len(values))
	for i, value := range values {
		bValue := reflect.Indirect(reflect.ValueOf(&bs[i]))
		setBool(&bValue, value)
	}
	bsValue := reflect.ValueOf(bs)
	keyValue.Set(bsValue)
	return nil
}

// Returns true on "1" and "true", anything returns false.
func setBool(keyValue *reflect.Value, value string) {
	var b bool
	value = strings.TrimSpace(value)
	if value == "1" || strings.ToLower(value) == "true" {
		b = true
	}
	keyValue.SetBool(b)
}

func setInts(keyValue *reflect.Value, values []string) error {
	var ns = make([]int, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setInt(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setInt8s(keyValue *reflect.Value, values []string) error {
	var ns = make([]int8, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setInt(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setInt16s(keyValue *reflect.Value, values []string) error {
	var ns = make([]int16, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setInt(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setInt32s(keyValue *reflect.Value, values []string) error {
	var ns = make([]int32, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setInt(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setInt64s(keyValue *reflect.Value, values []string) error {
	var ns = make([]int64, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setInt(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setInt(keyValue *reflect.Value, value string) error {
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return createConvertError(value, keyValue.Kind().String())
	}
	n64 := int64(n)

	if keyValue.OverflowInt(n64) {
		return createOverflowError(value, keyValue.Kind().String())
	}

	keyValue.SetInt(n64)
	return nil
}

func setUints(keyValue *reflect.Value, values []string) error {
	var ns = make([]uint, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setUint(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setUint8s(keyValue *reflect.Value, values []string) error {
	var ns = make([]uint8, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setUint(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setUint16s(keyValue *reflect.Value, values []string) error {
	var ns = make([]uint16, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setUint(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setUint32s(keyValue *reflect.Value, values []string) error {
	var ns = make([]uint32, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setUint(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setUint64s(keyValue *reflect.Value, values []string) error {
	var ns = make([]uint64, len(values))
	for i, value := range values {
		nValue := reflect.Indirect(reflect.ValueOf(&ns[i]))
		if err := setUint(&nValue, value); err != nil {
			return err
		}
	}
	nsValue := reflect.ValueOf(ns)
	keyValue.Set(nsValue)
	return nil
}

func setUint(keyValue *reflect.Value, value string) error {
	n, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return createConvertError(value, keyValue.Kind().String())
	}
	nu64 := uint64(n)

	if keyValue.OverflowUint(nu64) {
		return createOverflowError(value, keyValue.Kind().String())
	}

	keyValue.SetUint(nu64)
	return nil
}

func setFloat32s(keyValue *reflect.Value, values []string) error {
	var fs = make([]float32, len(values))
	for i, value := range values {
		fValue := reflect.Indirect(reflect.ValueOf(&fs[i]))
		if err := setFloat(&fValue, value); err != nil {
			return err
		}
	}
	fsValue := reflect.ValueOf(fs)
	keyValue.Set(fsValue)
	return nil
}

func setFloat64s(keyValue *reflect.Value, values []string) error {
	var fs = make([]float64, len(values))
	for i, value := range values {
		fValue := reflect.Indirect(reflect.ValueOf(&fs[i]))
		if err := setFloat(&fValue, value); err != nil {
			return err
		}
	}
	fsValue := reflect.ValueOf(fs)
	keyValue.Set(fsValue)
	return nil
}

func setFloat(keyValue *reflect.Value, value string) error {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return createConvertError(value, keyValue.Kind().String())
	}

	if keyValue.OverflowFloat(f) {
		return createOverflowError(value, keyValue.Kind().String())
	}

	keyValue.SetFloat(f)
	return nil
}

func setDurations(keyValue *reflect.Value, values []string) error {
	var ds = make([]time.Duration, len(values))
	for i, value := range values {
		dValue := reflect.Indirect(reflect.ValueOf(&ds[i]))
		if err := setDuration(&dValue, value); err != nil {
			return err
		}
	}
	dsValue := reflect.ValueOf(ds)
	keyValue.Set(dsValue)
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

func setTimes(keyValue *reflect.Value, values []string) error {
	var ts = make([]time.Time, len(values))
	for i, value := range values {
		tValue := reflect.Indirect(reflect.ValueOf(&ts[i]))
		if err := setTime(&tValue, value); err != nil {
			return err
		}
	}
	tsValue := reflect.ValueOf(ts)
	keyValue.Set(tsValue)
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
