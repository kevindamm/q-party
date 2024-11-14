// Copyright (c) 2024 Kevin Damm
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// github:kevindamm/q-party/dollar_value.go

package qparty

import (
	"errors"
	"fmt"
	"strconv"
)

type DollarValue int

func (value DollarValue) String() string {
	return fmt.Sprintf("$%d", int(value))
}

// Returns the integer value of a DollarValue, ignoring whether it was a wager.
func (value DollarValue) Abs() int {
	if value < 0 {
		return -int(value)
	}
	return int(value)
}

func (value DollarValue) IsWager() bool { return int(value) < 0 }
func (value *DollarValue) ToggleWager(next bool) {
	wager := int(*value) < 0
	if wager != next {
		*value = -*value
	}
}

func (value *DollarValue) UnmarshalText(text []byte) error {
	var err error
	*value, err = ParseDollarValue(string(text))
	return err
}

const ZeroDollars = DollarValue(0)

func ParseDollarValue(text string) (DollarValue, error) {
	if len(text) == 0 {
		return ZeroDollars, errors.New("empty string for dollar value")
	}
	if text[0] != '$' {
		return ZeroDollars, fmt.Errorf("invalid DollarValue %s, expected '$'", text)
	}

	intVal, err := strconv.Atoi(string(text[1:]))
	if err != nil {
		return ZeroDollars, fmt.Errorf("failed to parse integer %s\n%s", text[1:], err)
	}
	return DollarValue(intVal), nil
}
