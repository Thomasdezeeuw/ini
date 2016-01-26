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

// AppName is the name of our application.
var AppName string

func main() {
	// Open our configuration file.
	f, err := os.Open("./config.ini")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Parse the configuration.
	config, err := ini.Parse(f)
	if err != nil {
		panic(err)
	}

	// We can now access the configuration options like so.
	AppName = config[ini.Global]["Name"]
	address := config["HTTP"]["Host"] + ":" + config["HTTP"]["Port"]

	// AppName will be "My app" and address is "example.com:80".

	log.Println("Starting application: " + config[ini.Global]["Name"])
	log.Println("Listening on address " + address)

	err = http.ListenAndServe(address, http.HandlerFunc(homeHandler))
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Writes "Hello from My app".
	w.Write([]byte("Hello from " + AppName))
}
