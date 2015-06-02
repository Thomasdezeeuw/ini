package main

import (
	"fmt"
	"os"

	"github.com/Thomasdezeeuw/ini"
)

func main() {
	f, err := os.Open("config.ini")
	if err != nil {
		panic(err)
	}

	// Parse the configuration.
	c, err := ini.Parse(f)
	if err != nil {
		panic(err)
	}

	// We can now access it like so:
	fmt.Println(c["HTTP"]["Url"])  // localhost (string)
	fmt.Println(c["HTTP"]["Port"]) // 8000 (string)

	// Bud usually we want to use our own custom configuration, like so:
	var config struct {
		Name string
		HTTP struct {
			Url  string
			Port int
		}
	}

	// Now we can scan the ini.Config into our own configuration.
	c.Scan(&config)

	// And access its values like so:
	fmt.Println(config.HTTP.Url)  // localhost (string)
	fmt.Println(config.HTTP.Port) // 8000 (int)
}
