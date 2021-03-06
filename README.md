# Ini

[![GoDoc](https://godoc.org/github.com/Thomasdezeeuw/ini?status.svg)](https://godoc.org/github.com/Thomasdezeeuw/ini)
[![Build Status](https://img.shields.io/travis/Thomasdezeeuw/ini.svg)](https://travis-ci.org/Thomasdezeeuw/ini)
[![Coverage Status](https://coveralls.io/repos/Thomasdezeeuw/ini/badge.svg?branch=master&service=github)](https://coveralls.io/github/Thomasdezeeuw/ini?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Thomasdezeeuw/ini)](https://goreportcard.com/report/github.com/Thomasdezeeuw/ini)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Thomasdezeeuw/ini/blob/master/LICENSE)

Ini is a [Go](https://golang.org/) package for for parsing ini (or cfg) files.
See [Godoc](https://godoc.org/github.com/Thomasdezeeuw/ini) for the API.

## Stable

The api is now stable and will remain stable until version 2 is released. To
prove how stable `ini.Parse` is, it battled [go-fuzz](https://github.com/dvyukov/go-fuzz)
for 10 hours, processing over 1.1 billion randomized inputs and won:

```
2015/07/21 23:58:54 slaves: 8, corpus: 611 (3h8m ago), crashers: 0, restarts: 1/
10000, execs: 1100121441 (30462/sec), cover: 493, uptime: 10h1m
```

## Installation

Run the following line to install.

```bash
$ go get github.com/Thomasdezeeuw/ini
```

## Examples

See the [_examples](https://github.com/Thomasdezeeuw/ini/tree/master/_examples)
directory for a multiple examples.

## License

Licensed under the MIT license, copyright (C) Thomas de Zeeuw. See `LICENSE` for
more information.
