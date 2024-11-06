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
// github:kevindamm/q-party/cmd/jarchive/show_date.go

package main

import (
	"fmt"
	"strconv"
	"time"
)

type ShowDateRange struct {
	From ShowDate `json:"from,omitempty"`
	To   ShowDate `json:"until,omitempty"`
}

type ShowDate time.Time

var unknown_airing = ShowDate(time.Time{})
var unknown_taping = ShowDate(time.Time{})

func (date ShowDate) String() string {
	return fmt.Sprintf("%4d/%02d/%02d", date.Year(), date.Month(), date.MonthDay())
}

func (date ShowDate) Year() int {
	return time.Time(date).Year()
}

// Returns 0 if unknown, 1 for January, ..., 12 for December.
func (date ShowDate) Month() time.Month {
	return time.Time(date).Month()
}

func (date ShowDate) MonthDay() int {
	return time.Time(date).Day()
}

// Returns 0 for Sunday, 1 for Monday, ..., 6 for Saturday.
func (date ShowDate) WeekDay() time.Weekday {
	return time.Time(date).Weekday()
}

func (date ShowDate) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf(
			"%4d/%02d/%02d",
			date.Year(), int(date.Month()), date.MonthDay())),
		nil
}

func (date *ShowDate) UnmarshalText(text []byte) error {
	if len(text) != len("YYYY_MM_DD") {
		return fmt.Errorf("incorrect format for aired date '%s'", text)
	}

	year, err := strconv.Atoi(string(text[:4]))
	if err != nil {
		return err
	}

	month, err := strconv.Atoi(string(text[5:7]))
	if err != nil {
		return err
	}

	day, err := strconv.Atoi(string(text[8:]))
	if err != nil {
		return err
	}

	*date = ShowDate(time.Date(
		year, time.Month(month), day, 23, 0, 0, 0, time.UTC))
	return nil
}
