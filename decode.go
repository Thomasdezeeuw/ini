// Copyright (C) 2015-2016 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package ini

import (
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var timeFormats = []string{"2006-01-02", "2006-01-02 15:04",
	"2006-01-02 15:04:05", time.RFC3339, time.RFC1123, time.RFC822}

var (
	kindString  = reflect.TypeOf("").Kind()
	kindBool    = reflect.TypeOf(true).Kind()
	kindInt     = reflect.TypeOf(int(1)).Kind()
	kindInt8    = reflect.TypeOf(int8(1)).Kind()
	kindInt16   = reflect.TypeOf(int16(1)).Kind()
	kindInt32   = reflect.TypeOf(int32(1)).Kind()
	kindInt64   = reflect.TypeOf(int64(1)).Kind()
	kindUint    = reflect.TypeOf(uint(1)).Kind()
	kindUint8   = reflect.TypeOf(uint8(1)).Kind()
	kindUint16  = reflect.TypeOf(uint16(1)).Kind()
	kindUint32  = reflect.TypeOf(uint32(1)).Kind()
	kindUint64  = reflect.TypeOf(uint64(1)).Kind()
	kindFloat32 = reflect.TypeOf(float32(1.0)).Kind()
	kindFloat64 = reflect.TypeOf(float64(1.0)).Kind()

	typeDuration = reflect.TypeOf(time.Nanosecond)
	typeTime     = reflect.TypeOf(time.Time{})
)

// Decode decodes a configuration into a struct or map, see `Config.Decode`.
func Decode(r io.Reader, dst interface{}) error {
	c, err := Parse(r)
	if err != nil {
		return err
	}
	return c.Decode(dst)
}

// DecodeValue decodes a single configuration value into a variable.
func DecodeValue(value string, dst interface{}) error {
	valuePtr := reflect.ValueOf(dst)
	v := reflect.Indirect(valuePtr)

	// If it's not a pointer we can't change the original value.
	if valuePtr.Kind() != reflect.Ptr {
		return errors.New("ini: DecodeValue requires a pointer to a destination value")
	} else if !v.IsValid() || !v.CanSet() {
		return errors.New("ini: can't change value of destination value")
	}

	return setReflectValue(&v, value)
}

func setReflectValue(keyValue *reflect.Value, value string) error {
	if keyValue.Kind() == reflect.Slice {
		return setSlice(keyValue, value)
	}

	// Special time cases.
	switch keyValue.Type() {
	case typeDuration:
		return setDuration(keyValue, value)
	case typeTime:
		return setTime(keyValue, value)
	}

	switch keyValue.Kind() {
	case kindString:
		keyValue.SetString(value)
	case kindBool:
		return setBool(keyValue, value)
	case kindInt, kindInt8, kindInt16, kindInt32, kindInt64:
		return setInt(keyValue, value)
	case kindUint, kindUint8, kindUint16, kindUint32, kindUint64:
		return setUint(keyValue, value)
	case kindFloat32, kindFloat64:
		return setFloat(keyValue, value)
	}

	return nil
}

// SetSlice sets a slice of reflected value.
func setSlice(keyValue *reflect.Value, value string) error {
	values := getValues(value)

	// Special time cases.
	switch keyValue.Type().Elem() {
	case typeDuration:
		return setDurations(keyValue, values)
	case typeTime:
		return setTimes(keyValue, values)
	}

	// Type of the slice elements.
	switch keyValue.Type().Elem().Kind() {
	case kindString:
		keyValue.Set(reflect.ValueOf(values))
	case kindBool:
		return setBools(keyValue, values)
	case kindInt:
		return setInts(keyValue, values)
	case kindInt8:
		return setInt8s(keyValue, values)
	case kindInt16:
		return setInt16s(keyValue, values)
	case kindInt32:
		return setInt32s(keyValue, values)
	case kindInt64:
		return setInt64s(keyValue, values)
	case kindUint:
		return setUints(keyValue, values)
	case kindUint8:
		return setUint8s(keyValue, values)
	case kindUint16:
		return setUint16s(keyValue, values)
	case kindUint32:
		return setUint32s(keyValue, values)
	case kindUint64:
		return setUint64s(keyValue, values)
	case kindFloat32:
		return setFloat32s(keyValue, values)
	case kindFloat64:
		return setFloat64s(keyValue, values)
	}

	return nil
}

func setBools(keyValue *reflect.Value, values []string) error {
	var bs = make([]bool, len(values))
	for i, value := range values {
		bValue := reflect.Indirect(reflect.ValueOf(&bs[i]))
		if err := setBool(&bValue, value); err != nil {
			return err
		}
	}
	bsValue := reflect.ValueOf(bs)
	keyValue.Set(bsValue)
	return nil
}

func setBool(keyValue *reflect.Value, value string) error {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return createCovertionError(value, keyValue.Kind().String())
	}

	keyValue.SetBool(b)
	return nil
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
		return createCovertionError(value, keyValue.Kind().String())
	}
	n64 := n

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
		return createCovertionError(value, keyValue.Kind().String())
	}
	nu64 := n

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
		return createCovertionError(value, keyValue.Kind().String())
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
		return createCovertionError(value, "time.Duration")
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

	return createCovertionError(value, "time.Time")
}

func getValues(value string) []string {
	values := strings.Split(value, ",")
	for i, value := range values {
		values[i] = strings.TrimSpace(value)
	}
	return values
}
