// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

// todo: test int/uint/float overflows.

type completeTestData struct {
	Str      string
	Duration time.Duration
	Time     time.Time
	Bools    Bools
	Ints     Ints
	Uints    Uints
	Floats   Floats
	Slices   Slices
}

type Bools struct {
	T1 bool
	T2 bool
	T3 bool
	F1 bool
	F2 bool
	F3 bool
	F4 bool
}

type Ints struct {
	I   int
	I8  int8
	I16 int16
	I32 int32
	I64 int64
}

type Uints struct {
	Ui   uint
	Ui8  uint8
	Ui16 uint16
	Ui32 uint32
	Ui64 uint64
}

type Floats struct {
	F32 float32
	F64 float64
}

type Slices struct {
	Str      []string
	Duration []time.Duration
	Time     []time.Time
	T        []bool
	F        []bool
	I        []int
	I8       []int8
	I16      []int16
	I32      []int32
	I64      []int64
	Ui       []uint
	Ui8      []uint8
	Ui16     []uint16
	Ui32     []uint32
	Ui64     []uint64
	F32      []float32
	F64      []float64
}

func TestConfigDecode(t *testing.T) {
	t.Parallel()
	c := Config{
		Global: {
			"str":      "string",
			"duration": "5s",
			"time":     "2015-05-08 11:07:30",
		},
		"bools": {
			"t1": "true",
			"t2": "TRUE",
			"t3": "1",
			"f1": "false",
			"f2": "FALSE",
			"f3": "0",
			"f4": "anything else",
		},
		"ints": {
			"i":   "0",
			"i8":  "127",
			"i16": "32767",
			"i32": "2147483647",
			"i64": "9223372036854775807",
		},
		"uints": {
			"ui":   "0",
			"ui8":  "255",
			"ui16": "65535",
			"ui32": "4294967295",
			"ui64": "18446744073709551615",
		},
		"floats": {
			"f32": "12.21",
			"f64": "123.321",
		},
		"slices": {
			"str":      "string1, string2",
			"duration": "10s, 20m",
			"time":     "2015-05-10, 2015-05-09 11:07, 2015-05-08 11:07:30",
			"t":        "true, TRUE, 1",
			"f":        "false, FALSE, 0, anything else",
			"i":        "1, 2, 3",
			"i8":       "124, -125, 126",
			"i16":      "2343, 123, -23423",
			"i32":      "3534534, -234234, 86767",
			"i64":      "53534534530, 65756756, 4365456",
			"ui":       "1, 23425",
			"ui8":      "154, 126",
			"ui16":     "4645, 4353",
			"ui32":     "46424535, 3453",
			"ui64":     "3234464645, 453",
			"f32":      "12.21, 123.21",
			"f64":      "123.321, 4234.123",
		},
	}

	var got completeTestData

	if err := c.Decode(&got); err != nil {
		t.Fatalf("Unexpected error decoding config into variable: %q", err.Error())
	}

	strTimes := []string{"2015-05-10", "2015-05-09 11:07", "2015-05-08 11:07:30"}
	var times []time.Time
	for i, tValue := range strTimes {
		t1, err := time.Parse(timeFormats[i], tValue)
		if err != nil {
			t.Fatalf("Unexpected error parsing time: %s", err.Error())
		}
		times = append(times, t1)
	}

	expected := completeTestData{
		Str:      "string",
		Duration: 5 * time.Second,
		Time:     times[2],
		Bools: Bools{
			T1: true,
			T2: true,
			T3: true,
			F1: false,
			F2: false,
			F3: false,
			F4: false,
		},
		Ints: Ints{
			I:   0,
			I8:  127,
			I16: 32767,
			I32: 2147483647,
			I64: 9223372036854775807,
		},
		Uints: Uints{
			Ui:   0,
			Ui8:  255,
			Ui16: 65535,
			Ui32: 4294967295,
			Ui64: 18446744073709551615,
		},
		Floats: Floats{
			F32: 12.21,
			F64: 123.321,
		},
		Slices: Slices{
			Str:      []string{"string1", "string2"},
			Duration: []time.Duration{10 * time.Second, 20 * time.Minute},
			Time:     times,
			T:        []bool{true, true, true},
			F:        []bool{false, false, false, false},
			I:        []int{1, 2, 3},
			I8:       []int8{124, -125, 126},
			I16:      []int16{2343, 123, -23423},
			I32:      []int32{3534534, -234234, 86767},
			I64:      []int64{53534534530, 65756756, 4365456},
			Ui:       []uint{1, 23425},
			Ui8:      []uint8{154, 126},
			Ui16:     []uint16{4645, 4353},
			Ui32:     []uint32{46424535, 3453},
			Ui64:     []uint64{3234464645, 453},
			F32:      []float32{12.21, 123.21},
			F64:      []float64{123.321, 4234.123},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("Expected Config.Decode() to return %v, but got %v", expected, got)
	}
}

type smallTestData struct {
	Key     string
	Section smallSection
}

type smallSection struct {
	Key2 string
}

func TestDecode(t *testing.T) {
	t.Parallel()
	content := "key = value\n[section]\nkey2=value2"

	var got smallTestData
	if err := Decode(strings.NewReader(content), &got); err != nil {
		t.Fatalf("Unexpected error decoding: %q", err.Error())
	}

	expected := smallTestData{
		Key: "value",
		Section: smallSection{
			Key2: "value2",
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("Expected %v, but got %v", expected, got)
	}
}

func TestDecodeValue(t *testing.T) {
	t.Parallel()

	var host string
	var port int
	conf := Config{
		Global: {
			"host": "localhost",
			"port": "8080",
		},
	}

	if err := DecodeValue(conf[Global]["host"], &host); err != nil {
		t.Fatalf("Unexpected error decoding variable: %s", err.Error())
	} else if err := DecodeValue(conf[Global]["port"], &port); err != nil {
		t.Fatalf("Unexpected error decoding variable: %s", err.Error())
	}

	if expected := "localhost"; host != expected {
		t.Fatalf("Expected host to be %q, but got %q", expected, host)
	}
	if expected := 8080; port != expected {
		t.Fatalf("Expected port to be %d, but got %d", expected, port)
	}
}

func TestDecodeValueNotPointerError(t *testing.T) {
	t.Parallel()

	var ui8 uint8
	conf := Config{
		Global: {
			"ui8": "500",
		},
	}

	err := DecodeValue(conf[Global]["ui8"], ui8)
	expected := "ini: DecodeValue requires a pointer to a destination value"
	if err == nil {
		t.Fatal("Expected an error, but didn't get one")
	} else if err.Error() != expected {
		t.Fatalf("Expected error message to be %q, but got %q",
			expected, err.Error())
	}
}

func TestDecodeValueOverflowError(t *testing.T) {
	t.Parallel()

	var ui8 uint8
	conf := Config{
		Global: {
			"ui8": "500",
		},
	}

	err := DecodeValue(conf[Global]["ui8"], &ui8)
	expected := "ini: can't convert '500' to type uint8, it overflows the type"
	if err == nil {
		t.Fatal("Expected an error, but didn't get one")
	} else if err.Error() != expected {
		t.Fatalf("Expected error message to be %q, but got %q",
			expected, err.Error())
	} else if !IsOverflowError(err) {
		t.Fatal("Expected error to be an overflow error, but it isn't")
	}
}
