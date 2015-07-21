// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed onder the MIT license that can be found in the LICENSE file.

package ini

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"unicode"
)

const (
	commentStart byte = ';'
	seperator    byte = '='
	sectionStart byte = '['
	sectionEnd   byte = ']'
	escape       byte = '\\'
	lineEnd      byte = '\n'
	doubleQuote  byte = '"'
	singleQuote  byte = '\''
	nilQuote     byte = 0
)

type parser struct {
	Config            Config
	scanner           *bufio.Scanner
	currentSection    string
	currentlineNumber int
}

func (p *parser) parse() error {
	for p.scanner.Scan() {
		line := p.scanner.Bytes()

		if err := p.handleLine(line); err != nil {
			return createSynthaxError(p.currentlineNumber, err.Error())
		}
	}

	if err := p.scanner.Err(); err != nil {
		return fmt.Errorf("ini: error reading: %s", err.Error())
	}
	return nil
}

func (p *parser) handleLine(line []byte) error {
	line = bytes.TrimSpace(line)
	p.currentlineNumber++

	if len(line) == 0 {
		return nil
	}

	switch line[0] {
	case commentStart:
	case sectionStart:
		sectionName, err := parseSection(line)
		if err != nil {
			return err
		}

		if err := p.updateSection(sectionName); err != nil {
			return err
		}
	default:
		key, value, err := parseKeyValue(line)
		if err != nil {
			return err
		}

		if err := p.addKeyValue(key, value); err != nil {
			return err
		}
	}

	return nil
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
// *Note the reader already gets buffered, so there is no need to buffer it
// youself.
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

func parseSection(line []byte) (string, error) {
	// This should not be possible in regular execution, but just in case.
	if line[0] != sectionStart {
		return "", fmt.Errorf("section should start with %q", string(sectionStart))
	}

	l := len(line)
	if l > 2 {
		l -= 2
	}
	var sectionName = make([]byte, 0, l)
	var sectionEnded bool

	// Skipping the opening bracket.
	for i := 1; i < len(line); i++ {
		b := line[i]

		if b == sectionEnd {
			sectionEnded = true
			continue
		} else if b == commentStart && sectionEnded {
			break
		} else if sectionEnded && !unicode.IsSpace(rune(b)) {
			return "", fmt.Errorf("unexpected %q after section closed", string(b))
		}

		sectionName = append(sectionName, b)
	}

	if !sectionEnded {
		return "", errors.New("unclosed section")
	}

	section := string(bytes.TrimSpace(sectionName))
	if len(section) == 0 {
		return "", errors.New("section can't be empty")
	}

	return section, nil
}

func parseKeyValue(line []byte) (key, value string, err error) {
	var values [2][]byte
	var hasSeperator bool

	// Keep track of the byte number for both the key and value.
	for i, valueNumber := 0, 0; valueNumber <= 1; valueNumber++ {
		var isQouted, isEscaped, nextShouldBeSeperator bool
		var usedQouted byte

		for ; i < len(line); i++ {
			b := line[i]

			if nextShouldBeSeperator && !unicode.IsSpace(rune(b)) && b != seperator {
				return "", "", fmt.Errorf("unexpected %q, expected the seperator %q",
					string(b), string(seperator))
			} else if (b == doubleQuote || b == singleQuote) && !isEscaped {
				if !isQouted {
					isQouted = true
					usedQouted = b
					continue
				} else if b == usedQouted {
					isQouted = false
					usedQouted = nilQuote
					if valueNumber == 0 {
						nextShouldBeSeperator = true
					}
					continue
				}
			} else if b == commentStart && !isQouted && hasSeperator {
				break
			} else if b == escape && !isEscaped {
				isEscaped = true
				continue
			} else if b == seperator && valueNumber == 0 && !isQouted {
				hasSeperator = true
				i++ // skip the separator in the next loop.
				break
			}

			isEscaped = false
			values[valueNumber] = append(values[valueNumber], b)
		}

		if isQouted {
			return "", "", errors.New("qoute not closed")
		}
	}

	if !hasSeperator {
		return "", "", errors.New("no separator found")
	}

	key = string(bytes.TrimSpace(values[0]))
	value = string(bytes.TrimSpace(values[1]))
	if len(key) == 0 {
		return "", "", errors.New("key can't be empty")
	}
	return key, value, nil
}
