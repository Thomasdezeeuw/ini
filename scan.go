package ini

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Time formats supported in parsing time.
var timeFormats = []string{"2006-01-02 15:04", "2006-01-02 15:04:05"}

// SetReflectValue sets the value of reflected value.
func setReflectValue(f *reflect.Value, value string) error {
	// Test for time.Duration and time.Time first.
	// todo: improve this, this really shouldn't be the way to do it.
	if tStr := fmt.Sprintf("%v", f.Type()); tStr == "time.Duration" {
		duration, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		f.Set(reflect.ValueOf(duration))
		return nil
	} else if tStr == "time.Time" {
		for _, format := range timeFormats {
			t, err := time.Parse(format, value)
			if err == nil {
				f.Set(reflect.ValueOf(t))
				return nil
			}
		}
		return errors.New("ini: unkown time layout " + value)
	}

	// todo: clean up the switch.
	switch f.Kind() {
	case reflect.String:
		f.SetString(value)
	case reflect.Bool:
		f.SetBool(getBool(value))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		} else if f.OverflowInt(int64(i)) {
			return fmt.Errorf("ini: %d overflows %q", i, f.Kind().String())
		}

		f.SetInt(int64(i))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		} else if f.OverflowUint(uint64(i)) {
			return fmt.Errorf("ini: %d overflows %q", i, f.Kind().String())
		}

		f.SetUint(uint64(i))
	case reflect.Float32, reflect.Float64:
		fv, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		} else if f.OverflowFloat(fv) {
			return fmt.Errorf("ini: %f overflows %q", fv, f.Kind().String())
		}

		f.SetFloat(fv)
	case reflect.Slice:
		err := setSlice(f, value)
		if err != nil {
			return err
		}
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
			bs = append(bs, getBool(value))
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
func getBool(value string) bool {
	value = strings.TrimSpace(value)
	if value == "1" || strings.ToLower(value) == "true" {
		return true
	}
	return false
}
