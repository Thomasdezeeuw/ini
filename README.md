# Ini

[![GoDoc](https://godoc.org/github.com/Thomasdezeeuw/ini?status.svg)](https://godoc.org/github.com/Thomasdezeeuw/ini)
[![Build Status](https://travis-ci.org/Thomasdezeeuw/ini.png?branch=master)](https://travis-ci.org/Thomasdezeeuw/ini)

Ini is a [Go](https://golang.org/) package for for parsing ini (or cfg) files.

## Installation

Run the following line to install.

```bash
$ go get github.com/Thomasdezeeuw/ini
```

## Usage

You can load a configuration from a ini formatted file.

```go
package main

import (
	"fmt"

	"github.com/Thomasdezeeuw/ini"
)

func main() {
	c, err := ini.Load("development.ini") // or production
	if err != nil {
		fmt.Println(err)
		return
	}

	var config struct {
		Name string
		HTTP struct {
			Url  string
			Port int
		}
	}

	c.Scan(&config)

	fmt.Println(config.Name)       // My app
	fmt.Println(config.HTTP.Url)   // localhost
	fmt.Println(config.HTTP.Port)  // 8000 (int)
	fmt.Println(c["HTTP"]["Port"]) // 8000 (string)
}
```

```
; development.ini
Name = "My app"

[HTTP]
Url = "localhost"
Port = 8000
```

## License

Licensed under the MIT license, copyright (C) Thomas de Zeeuw.