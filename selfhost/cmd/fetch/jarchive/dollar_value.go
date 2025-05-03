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

package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/kevindamm/q-party/schema"
)

const ZeroDollars = schema.Value(0)

func ParseDollarValue(text string) (schema.Value, error) {
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
	return schema.Value(intVal), nil
}
