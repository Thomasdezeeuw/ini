package ini

import (
	"fmt"
	"strings"
)

var config Config

func ExampleGlobal() {
	value := config[Global]["key"]
	value2, found := config[Global]["key2"]

	fmt.Println("value 1:", value)
	fmt.Println("value 2 found:", found)
	fmt.Println("value 2:", value2)
}

func ExampleParse() {
	r := strings.NewReader("key = value")

	conf, err := Parse(r)
	if err != nil {
		panic(err)
	}

	value := conf[Global]["key"]
	fmt.Println("value:", value)

	value, found := conf[Global]["unkown"]
	fmt.Println("found:", found)
	fmt.Println("value:", value)
	// Output:
	// value: value
	// found: false
	// value:
}

func ExampleParse_section() {
	confFile := "key = value\n" +
		"[section1]\n" +
		"key = value2"
	r := strings.NewReader(confFile)

	conf, err := Parse(r)
	if err != nil {
		panic(err)
	}

	value := conf[Global]["key"]
	value2 := conf["section1"]["key"]
	fmt.Println("value:", value)
	fmt.Println("value2:", value2)
	// Output:
	// value: value
	// value2: value2
}
