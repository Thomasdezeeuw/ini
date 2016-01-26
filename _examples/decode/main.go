// Copyright (C) 2015-2016 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
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

// AppName is the name of our application.
var AppName string

func main() {
	config := getConfig("./config.ini")

	// We can now access the configuration options like so.
	// AppName will be "My app".
	AppName = config.Name

	log.Println("Starting application: " + AppName)
	log.Println("Listening on address " + config.HTTP.Address)

	err := http.ListenAndServe(config.HTTP.Address, http.HandlerFunc(homeHandler))
	if err != nil {
		panic(err)
	}
}

func getConfig(filepath string) config {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Parse the configuration.
	iniConfig, err := ini.Parse(f)
	if err != nil {
		panic(err)
	}

	// But we usually want to use our own custom configuration. We can decode our
	// configuration into our own configuration variable, like so.
	var conf config
	if err := iniConfig.Decode(&conf); err != nil {
		panic(err)
	}

	// But of course we can still use the ini configuration.
	conf.HTTP.Address = conf.HTTP.Host + ":" + iniConfig["HTTP"]["Port"]

	return conf
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Writes "Hello from My app".
	w.Write([]byte("Hello from " + AppName))
}
