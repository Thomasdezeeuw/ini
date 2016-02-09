// Copyright (C) 2015-2016 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package ini

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

const (
	commentStart1 byte = ';'
	commentStart2 byte = '#'
	separator     byte = '='
	sectionStart  byte = '['
	sectionEnd    byte = ']'
	escape        byte = '\\'
	doubleQuote   byte = '"'
	singleQuote   byte = '\''
	nilQuote      byte = 0
)

type parser struct {
	Config         Config
	scanner        *bufio.Scanner
	currentSection string
}

func (p *parser) parse() error {
	var currentlineNumber int

	for p.scanner.Scan() {
		line := p.scanner.Bytes()
		currentlineNumber++

		if err := p.handleLine(line); err != nil {
			return createSynthaxError(currentlineNumber, err.Error())
		}
	}

	if err := p.scanner.Err(); err != nil {
		return fmt.Errorf("ini: error reading: %s", err.Error())
	}
	return nil
}

func (p *parser) handleLine(line []byte) error {
	line = bytes.TrimSpace(line)
	if len(line) == 0 {
		return nil
	}

	b := line[0]
	if isCommentStart(b) {
		return nil
	} else if b == sectionStart {
		sectionName, err := parseSection(line)
		if err != nil {
			return err
		}
		return p.updateSection(sectionName)
	}

	key, value, err := parseKeyValue(line)
	if err != nil {
		return err
	}
	return p.addKeyValue(key, value)
}

func (p *parser) updateSection(sectionName string) error {
	if _, ok := p.Config[sectionName]; ok {
		return fmt.Errorf("section %q already exists", sectionName)
	}
	p.currentSection = sectionName
	p.Config[sectionName] = map[string]string{}
	return nil
}

func (p *parser) addKeyValue(key, value string) error {
	sectionName := p.currentSection
	if _, ok := p.Config[sectionName][key]; ok {
		if sectionName == Global {
			sectionName = globalName
		}
		return fmt.Errorf("key %q already used in section %q", key, sectionName)
	}

	p.Config[sectionName][key] = value
	return nil
}

// Parse parses ini formatted input.
//
// Note: the reader already gets buffered, so there is no need to buffer it
// yourself.
func Parse(r io.Reader) (Config, error) {
	p := newParser(r)
	if err := p.parse(); err != nil {
		return nil, err
	}

	return p.Config, nil
}

func newParser(r io.Reader) *parser {
	return &parser{
		Config:         Config{Global: {}},
		currentSection: Global,
		scanner:        bufio.NewScanner(r),
	}
}

// Assumes the first character is always an opening bracket.
func parseSection(line []byte) (string, error) {
	var end int
	var sectionEnded bool

	// Skipping the opening bracket.
	for i, l := 1, len(line); i < l; i++ {
		b := line[i]

		if b == sectionEnd {
			sectionEnded = true
			end = i
			continue
		} else if isCommentStart(b) && sectionEnded {
			break
		} else if sectionEnded && !unicode.IsSpace(rune(b)) {
			return "", fmt.Errorf("unexpected %q after section closed",
				getFullRune(line[i:]))
		}
	}

	if !sectionEnded {
		return "", errors.New("unclosed section")
	}

	section := string(bytes.TrimSpace(line[1:end]))
	if len(section) == 0 {
		return "", errors.New("section can't be empty")
	}

	return section, nil
}

func parseKeyValue(line []byte) (key, value string, err error) {
	var values [2][]byte
	var areQuoted [2]bool
	var hasSeparator bool

	// Keep track of the byte number for both the key and value.
	for i, valueNumber := 0, 0; valueNumber <= 1; valueNumber++ {
		var isQuoted, isEscaped, nextShouldBeSeparator bool
		var usedQuote byte

		for ; i < len(line); i++ {
			b := line[i]
			isSpace := unicode.IsSpace(rune(b))

			if areQuoted[valueNumber] && !isQuoted && isSpace {
				// Quoted value with whitespace after the closing quote.
				continue
			} else if nextShouldBeSeparator && !isSpace && b != separator {
				return "", "", fmt.Errorf("unexpected %q, expected the separator %q",
					getFullRune(line[i:]), string(separator))
			} else if (b == doubleQuote || b == singleQuote) && !isEscaped {
				if !isQuoted {
					isQuoted = true
					areQuoted[valueNumber] = true
					values[valueNumber] = []byte{}
					usedQuote = b
					continue
				} else if b == usedQuote {
					isQuoted = false
					usedQuote = nilQuote
					if valueNumber == 0 {
						nextShouldBeSeparator = true
					}
					continue
				}
			} else if isCommentStart(b) && !isQuoted && hasSeparator {
				break
			} else if b == escape && !isEscaped {
				isEscaped = true
				continue
			} else if b == separator && valueNumber == 0 && !isQuoted {
				hasSeparator = true
				i++ // skip the separator in the next loop.
				break
			}

			isEscaped = false
			values[valueNumber] = append(values[valueNumber], b)
		}

		if isQuoted {
			return "", "", errors.New("quote not closed")
		}
	}

	if !hasSeparator {
		return "", "", errors.New("no separator found")
	}

	// Only trim extra whitespace if the value weren't quoted.
	for i := 0; i < len(values); i++ {
		if !areQuoted[i] {
			values[i] = bytes.TrimSpace(values[i])
		}
	}

	key = string(values[0])
	value = string(values[1])
	if len(key) == 0 {
		return "", "", errors.New("key can't be empty")
	}
	return key, value, nil
}

func getFullRune(line []byte) string {
	r, _ := utf8.DecodeRune(line)
	return string(r)
}

func isCommentStart(b byte) bool {
	return b == commentStart1 || b == commentStart2
}
