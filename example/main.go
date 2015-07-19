package main

import (
	"fmt"
	"os"

	"github.com/Thomasdezeeuw/ini"
)

type config struct {
	Name string
	HTTP struct {
		Host    string
		Port    int
		Address string
	}
}

func main() {
	conf := getConfig("./config.ini")

	fmt.Println("Starting application:", conf.Name) // My app

	// And access its values like so:
	fmt.Println(conf.HTTP.Host)    // example.com
	fmt.Println(conf.HTTP.Port)    // 80
	fmt.Println(conf.HTTP.Address) // example.com:80
}

func getConfig(filepath string) config {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Parse the configuration.
	c, err := ini.Parse(f)
	if err != nil {
		panic(err)
	}

	// We can now access the configuration options like so:
	fmt.Println(c["HTTP"]["Host"]) // example.com
	fmt.Println(c["HTTP"]["Port"]) // 80
	fmt.Println()

	// But we usually want to use our own custom configuration. We can scan our
	// configuration into our own configuration variable, like so:
	var conf config
	c.Scan(&conf)

	// And then set another configuration option.
	conf.HTTP.Address = c["HTTP"]["Host"] + ":" + c["HTTP"]["Port"]

	return conf
}
