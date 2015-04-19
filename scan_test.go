package ini

import (
	"reflect"
	"testing"
	"time"
)

func TestGetBool(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"0", false},
		{"false", false},
		{"1", true},
		{"true", true},
		{"bad", false},
	}

	for _, test := range tests {
		b := getBool(test.input)
		if b != test.expected {
			t.Fatalf("Expected getBool(%q) to return %t, got %t",
				test.input, test.expected, b)
		}
	}
}

// todo: combine the TestSetSlice* functions.
// todo: test slice errors.
func TestSetSliceString(t *testing.T) {
	var got []string
	expected := []string{"tag1", "tag2"}
	input := "tag1, tag2"

	f := reflect.Indirect(reflect.ValueOf(&got))
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceBool(t *testing.T) {
	var got []bool
	var expected = []bool{false, false, true, true}
	input := "0, false, 1, true"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceInt(t *testing.T) {
	var got []int
	var expected = []int{-1, 0, 1, 10, 100}
	input := "-1, 0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceInt8(t *testing.T) {
	var got []int8
	var expected = []int8{-1, 0, 1, 10, 100}
	input := "-1, 0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceInt16(t *testing.T) {
	var got []int16
	var expected = []int16{-1, 0, 1, 10, 100}
	input := "-1, 0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceInt32(t *testing.T) {
	var got []int32
	var expected = []int32{-1, 0, 1, 10, 100}
	input := "-1, 0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceInt64(t *testing.T) {
	var got []int64
	var expected = []int64{-1, 0, 1, 10, 100}
	input := "-1, 0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceUint(t *testing.T) {
	var got []uint
	var expected = []uint{0, 1, 10, 100}
	input := "0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceUint8(t *testing.T) {
	var got []uint8
	var expected = []uint8{0, 1, 10, 100}
	input := "0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceUint16(t *testing.T) {
	var got []uint16
	var expected = []uint16{0, 1, 10, 100}
	input := "0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceUint32(t *testing.T) {
	var got []uint32
	var expected = []uint32{0, 1, 10, 100}
	input := "0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceUint64(t *testing.T) {
	var got []uint64
	var expected = []uint64{0, 1, 10, 100}
	input := "0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceUintptr(t *testing.T) {
	var got []uintptr
	var expected = []uintptr{0, 1, 10, 100}
	input := "0, 1, 10, 100"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceFloat32(t *testing.T) {
	var got []float32
	var expected = []float32{-1.0, 0.0, 1.0, 10.0, 100.0}
	input := "-1.0, 0.0, 1.0, 10.0, 100.0"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestSetSliceFloat64(t *testing.T) {
	var got []float64
	var expected = []float64{-1.0, 0.0, 1.0, 10.0, 100.0}
	input := "-1.0, 0.0, 1.0, 10.0, 100.0"

	f := reflect.ValueOf(&got)
	f = reflect.Indirect(f)
	if err := setSlice(&f, input); err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
			got, input, got, expected)
	}

	for i, g := range got {
		if e := expected[i]; g != e {
			t.Fatalf("Expected setSlice(%v, %q) to return %v, got %v",
				got, input, got, expected)
		}
	}
}

func TestConfigScan(t *testing.T) {
	var got struct {
		True     bool // "1"
		True2    bool // "true"
		False    bool // "0"
		False2   bool // "false"
		Float32  float32
		Float64  float64
		String   string
		Time     time.Time
		Duration time.Duration
		Slice    []string
		Ints     struct {
			Int   int
			Int8  int8
			Int16 int16
			Int32 int32
			Int64 int64
		}
		Uints struct {
			Uint    uint
			Uint8   uint8
			Uint16  uint16
			Uint32  uint32
			Uint64  uint64
			Uintptr uintptr
		}
	}

	c := Config{
		Global: {
			"NotInStruct": ":(",
			"True":        "1",
			"True2":       "true",
			"False":       "0",
			"False2":      "false",
			"Float32":     "1.00",
			"Float64":     "2.00",
			"String":      "string",
			"Time":        "2015-04-05 17:47",
			"Duration":    "5s",
			"Slice":       "string1, string2",
		},
		"Ints": {
			"Int":   "-1",
			"Int8":  "10",
			"Int16": "-100",
			"Int32": "1000",
			"Int64": "-10000",
		},
		"Uints": {
			"Uint":   "1",
			"Uint8":  "10",
			"Uint16": "100",
			"Uint32": "1000",
			"Uint64": "10000",
		},
	}

	if err := c.Scan(&got); err != nil {
		t.Fatal(err)
	}

	// todo: improve these checks.. kind of repetitive
	if expected := true; got.True != expected {
		t.Fatalf("Expected got.True to be %t, got %t", expected, got.True)
	}

	if expected := true; got.True2 != expected {
		t.Fatalf("Expected got.True2 to be %t, got %t", expected, got.True2)
	}

	if expected := false; got.False != expected {
		t.Fatalf("Expected got.False to be %t, got %t", expected, got.False)
	}

	if expected := false; got.False2 != expected {
		t.Fatalf("Expected got.False2 to be %t, got %t", expected, got.False2)
	}

	if expected := float32(1.00); got.Float32 != expected {
		t.Fatalf("Expected got.Float32 to be %f, got %f", expected, got.Float32)
	}

	if expected := float64(2.00); got.Float64 != expected {
		t.Fatalf("Expected got.Float64 to be %f, got %f", expected, got.Float64)
	}

	if expected := "string"; got.String != expected {
		t.Fatalf("Expected got.String to be %q, got %q", expected, got.String)
	}

	if expected, _ := time.Parse("2006-01-02 15:04", "2015-04-05 17:47"); got.Time != expected {
		t.Fatalf("Expected got.Time to be %v, got %v", expected, got.Time)
	}

	if expected := 5 * time.Second; got.Duration != expected {
		t.Fatalf("Expected got.Duration to be %v, got %v", expected, got.Duration)
	}

	if expected := []string{"string1", "string2"}; got.Slice[0] != expected[0] || got.Slice[1] != expected[1] {
		t.Fatalf("Expected got.Slice to be %v, got %v", expected, got.Slice)
	}

	if expected := -1; got.Ints.Int != expected {
		t.Fatalf("Expected got.Ints.Int to be %d, got %d", expected, got.Ints.Int)
	}

	if expected := int8(10); got.Ints.Int8 != expected {
		t.Fatalf("Expected got.Ints.Int8 to be %d, got %d", expected, got.Ints.Int8)
	}

	if expected := int16(-100); got.Ints.Int16 != expected {
		t.Fatalf("Expected got.Ints.Int16 to be %d, got %d", expected, got.Ints.Int16)
	}

	if expected := int32(1000); got.Ints.Int32 != expected {
		t.Fatalf("Expected got.Ints.Int32 to be %d, got %d", expected, got.Ints.Int32)
	}

	if expected := int64(-10000); got.Ints.Int64 != expected {
		t.Fatalf("Expected got.Ints.Int64 to be %d, got %d", expected, got.Ints.Int64)
	}

	if expected := uint(1); got.Uints.Uint != expected {
		t.Fatalf("Expected got.Uints.Uint to be %d, got %d", expected, got.Uints.Uint)
	}

	if expected := uint8(10); got.Uints.Uint8 != expected {
		t.Fatalf("Expected got.Uints.Uint8 to be %d, got %d", expected, got.Uints.Uint8)
	}

	if expected := uint16(100); got.Uints.Uint16 != expected {
		t.Fatalf("Expected got.Uints.Uint16 to be %d, got %d", expected, got.Uints.Uint16)
	}

	if expected := uint32(1000); got.Uints.Uint32 != expected {
		t.Fatalf("Expected got.Uints.Uint32 to be %d, got %d", expected, got.Uints.Uint32)
	}

	if expected := uint64(10000); got.Uints.Uint64 != expected {
		t.Fatalf("Expected got.Uints.Uint64 to be %d, got %d", expected, got.Uints.Uint64)
	}
}

func TestScan(t *testing.T) {
	var dst struct {
		Name string
		Msg  string
		Http struct {
			Port int
			Url  string
		}
		Database struct {
			User     string
			Password string
		}
	}

	if err := Scan("testdata/config.ini", &dst); err != nil {
		t.Fatal(err)
	}

	if expected := "http server"; dst.Name != expected {
		t.Fatalf("Expected dst.Name to be %q, got %q", expected, dst.Name)
	}
	if expected := "Welcome \"Bob\""; dst.Msg != expected {
		t.Fatalf("Expected dst.Msg to be %q, got %q", expected, dst.Msg)
	}
	if expected := 8080; dst.Http.Port != expected {
		t.Fatalf("Expected dst.Http.Port to be %d, got %d", expected, dst.Http.Port)
	}
	if expected := "example.com"; dst.Http.Url != expected {
		t.Fatalf("Expected dst.Http.Url to be %q, got %q", expected, dst.Http.Url)
	}
	if expected := "bob"; dst.Database.User != expected {
		t.Fatalf("Expected dst.Database.User to be %q, got %q", expected, dst.Database.User)
	}
	if expected := "password"; dst.Database.Password != expected {
		t.Fatalf("Expected dst.Database.Password to be %q, got %q", expected, dst.Database.Password)
	}
}

// todo: combine the TestConfigScan*Error functions.
func TestConfigScanError(t *testing.T) {
	var got struct{}
	c := Config{}

	err := c.Scan(got)
	if err == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if expected := "ini: ini.Config.Scan requires a pointer to struct"; err.Error() != expected {
		t.Fatalf("Expected the error to be %q, but got %q", expected, err.Error())
	}
}

func TestConfigScanDurationError(t *testing.T) {
	var got struct {
		Duration time.Duration
	}

	c := Config{
		Global: {
			"Duration": ":(",
		},
	}

	err := c.Scan(&got)
	if err == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if expected := "time: invalid duration :("; err.Error() != expected {
		t.Fatalf("Expected the error to be %q, but got %q", expected, err.Error())
	}
}

func TestConfigScanTimeError(t *testing.T) {
	var got struct {
		Time time.Time
	}

	c := Config{
		Global: {
			"Time": "bad",
		},
	}

	err := c.Scan(&got)
	if err == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if expected := "ini: unkown time layout bad"; err.Error() != expected {
		t.Fatalf("Expected the error to be %q, but got %q", expected, err.Error())
	}
}

func TestConfigScanIntError(t *testing.T) {
	var got struct {
		Int int
	}

	c := Config{
		Global: {
			"Int": "bad",
		},
	}

	err := c.Scan(&got)
	if err == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if expected := "strconv.ParseInt: parsing \"bad\": invalid syntax"; err.Error() != expected {
		t.Fatalf("Expected the error to be %q, but got %q", expected, err.Error())
	}
}

func TestConfigScanUintError(t *testing.T) {
	var got struct {
		Uint uint
	}

	c := Config{
		Global: {
			"Uint": "bad",
		},
	}

	err := c.Scan(&got)
	if err == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if expected := "strconv.ParseInt: parsing \"bad\": invalid syntax"; err.Error() != expected {
		t.Fatalf("Expected the error to be %q, but got %q", expected, err.Error())
	}
}

func TestConfigScanFloatError(t *testing.T) {
	var got struct {
		Float float32
	}

	c := Config{
		Global: {
			"Float": "bad",
		},
	}

	err := c.Scan(&got)
	if err == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if expected := "strconv.ParseFloat: parsing \"bad\": invalid syntax"; err.Error() != expected {
		t.Fatalf("Expected the error to be %q, but got %q", expected, err.Error())
	}
}
